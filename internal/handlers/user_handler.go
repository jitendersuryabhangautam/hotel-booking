package handlers

import (
	"hotel-booking/internal/models"
	"hotel-booking/internal/response"
	"hotel-booking/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req models.RegisterRequest

	// Parse and validate JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		response.JSON(c, http.StatusBadRequest, false, "Invalid request", nil, err.Error())
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		response.JSON(c, http.StatusBadRequest, false, "All fields are required", nil, "Missing required fields")
		return
	}

	// 👉 Convert RegisterRequest ➔ User struct
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role, // 👈 Map this too
	}

	// Call service
	createdUser, err := h.service.RegisterUser(c.Request.Context(), user)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, false, "Failed to register", nil, err.Error())
		return
	}

	response.JSON(c, http.StatusCreated, true, "User registered successfully", createdUser, "")
}

func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")
	role, _ := c.Get("role")

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User info from token",
		"data": gin.H{
			"user_id": userID,
			"email":   email,
			"role":    role,
		},
	})
}

// internal/handlers/user_handler.go

func (h *UserHandler) GetMe(c *gin.Context) {
	// Extract the user from the context (set by middleware)
	userInterface, exists := c.Get("user")
	if !exists {
		response.JSON(c, http.StatusUnauthorized, false, "Unauthorized", nil, "No user found in context")
		return
	}

	user, ok := userInterface.(*models.User)
	if !ok {
		response.JSON(c, http.StatusInternalServerError, false, "Server error", nil, "Failed to parse user")
		return
	}

	// Success
	response.JSON(c, http.StatusOK, true, "User fetched successfully", user, "")
}
