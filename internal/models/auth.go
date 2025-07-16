package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User *User  `json:"user"`
	Token string `json:"token"`
}