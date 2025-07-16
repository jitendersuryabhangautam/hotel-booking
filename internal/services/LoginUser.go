package services

import (
	"context"
	"errors"
	"hotel-booking/internal/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// ✅ LOGIN FUNCTION
func (s *UserService) LoginUser(ctx context.Context, email, password string) (*models.LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil || user == nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := generateJWT(user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	user.Password = "" // 🛡️ Never return password

	return &models.LoginResponse{
		User:  user,
		Token: token,
	}, nil
}

// 🔑 JWT Token Generator
func generateJWT(user *models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "mysecret" // For dev only
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
