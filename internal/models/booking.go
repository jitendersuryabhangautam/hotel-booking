package models

import "time"

type Booking struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	RoomID         int       `json:"room_id"`
	CheckInDate    time.Time `json:"check_in_date"`
	CheckOutDate   time.Time `json:"check_out_date"`
	Adults         int       `json:"adults"`
	Children       int       `json:"children"`
	TotalAmount    float64   `json:"total_amount"`
	Status         string    `json:"status"`
	PaymentStatus  string    `json:"payment_status"`
	SpecialRequests string   `json:"special_requests"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateBookingRequest struct {
	UserID          int    `json:"user_id" binding:"required"`
	RoomID          int    `json:"room_id" binding:"required"`
	CheckInDate     string `json:"check_in_date" binding:"required,datetime=2006-01-02"`
	CheckOutDate    string `json:"check_out_date" binding:"required,datetime=2006-01-02"`
	Adults          int    `json:"adults" binding:"required,min=1,max=10"`
	Children        int    `json:"children" binding:"min=0,max=5"`
	SpecialRequests string `json:"special_requests" binding:"max=500"`
}

type BookingResponse struct {
	ID            int       `json:"id"`
	CheckInDate   time.Time `json:"check_in_date"`
	CheckOutDate  time.Time `json:"check_out_date"`
	Adults        int       `json:"adults"`
	Children      int       `json:"children"`
	TotalAmount   float64   `json:"total_amount"`
	Status        string    `json:"status"`
	PaymentStatus string    `json:"payment_status"`
	SpecialRequests *string  `json:"special_requests,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user"`

	Room struct {
		ID         int    `json:"id"`
		RoomNumber string `json:"room_number"`
		RoomType   string `json:"room_type"`
	} `json:"room"`
}



