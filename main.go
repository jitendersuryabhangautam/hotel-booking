package main

import (
	"fmt"
	"hotel-booking/internal/cache"
	"hotel-booking/internal/db"
	"hotel-booking/internal/handlers"
	"hotel-booking/internal/repositories"
	"hotel-booking/internal/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
			_ = godotenv.Load()

	// 2️⃣ Confirm env loaded
	fmt.Println("✅ DB_URL:", os.Getenv("DB_URL"))
	fmt.Println("✅ JWT_SECRET:", os.Getenv("JWT_SECRET"))

	// 3️⃣ Then init DB
	if err := db.Init(); err != nil {
		log.Fatalf("❌ Failed to initialize database: %v", err)
	}
	defer db.Close()
cache.Init()      // 👈 Redis Init
	defer cache.Close()
	// ✅ Confirm DB is not nil
	if db.DB == nil {
		log.Fatal("❌ Database connection is nil")
	}
	router := gin.Default()

	// User
	userRepo := repositories.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Room
	roomRepo := repositories.NewRoomRepository(db.DB)
	roomService := services.NewRoomService(roomRepo)
	roomHandler := handlers.NewRoomHandler(roomService)
	bookingRepo := repositories.NewBookingRepository(db.DB)
	bookingService := services.NewBookingService(bookingRepo)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)

		users := v1.Group("/users")
		users.GET("/me", userHandler.GetMe) // middleware can be added later
	
		bookings := v1.Group("/bookings")
		bookings.POST("", bookingHandler.CreateBooking)
		bookings.GET("/:id", bookingHandler.GetBookingByID)
		bookings.GET("", bookingHandler.GetAllBookings)
		rooms := v1.Group("/rooms")
		rooms.GET("/:id", roomHandler.GetRoomByID)	
		rooms.POST("/", roomHandler.CreateRoom)
		rooms.PUT("/:id", roomHandler.UpdateRoom)
		rooms.DELETE("/:id", roomHandler.DeleteRoom)
		rooms.GET("/available", roomHandler.GetAvailableRooms) // ✅ New API
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Server running on http://localhost:" + port)
	router.Run(":" + port)
}
