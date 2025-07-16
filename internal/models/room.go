package models

import "time"

type Room struct {
	ID             int       `json:"id"`
	RoomNumber     string    `json:"room_number"`
	RoomType       string    `json:"room_type"`
	Description    string    `json:"description"`
	PricePerNight  float64   `json:"price_per_night"`
	Capacity       int       `json:"capacity"`
	Floor          int       `json:"floor"`
	Amenities      []string  `json:"amenities"`
	IsAvailable    bool      `json:"is_available"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// type CreateRoomRequest struct {
// 	RoomNumber     string   `json:"room_number"`
// 	RoomType       string   `json:"room_type"`
// 	Description    string   `json:"description"`
// 	PricePerNight  float64  `json:"price_per_night"`
// 	Capacity       int      `json:"capacity"`
// 	Floor          int      `json:"floor"`
// 	Amenities      []string `json:"amenities"`
// 	IsAvailable    bool     `json:"is_available"`
// }
type CreateRoomRequest struct {
	RoomNumber    string  `json:"room_number" validate:"required"`
	RoomType      string  `json:"room_type" validate:"required,oneof=single double suite"`
	Description   string  `json:"description"`
	PricePerNight float64 `json:"price_per_night" validate:"required,gt=0"`
	Capacity      int     `json:"capacity" validate:"required,gt=0"`
	Floor         int     `json:"floor" validate:"required"`
	Amenities     []string  `json:"amenities"`
	IsAvailable   bool    `json:"is_available"`
}
type UpdateRoomRequest struct {
	RoomNumber     string   `json:"room_number"`
	RoomType       string   `json:"room_type"`
	Description    string   `json:"description"`
	PricePerNight  float64  `json:"price_per_night"`
	Capacity       int      `json:"capacity"`
	Floor          int      `json:"floor"`
	Amenities      []string `json:"amenities"`
	IsAvailable    bool     `json:"is_available"`
}