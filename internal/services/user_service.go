package services

import (
	"context"
	"errors"
	"hotel-booking/internal/models"
	"hotel-booking/internal/repositories"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
		return nil, errors.New("all fields required")
	}

	existing, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err == nil && existing != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	if user.Role == "" {
		user.Role = "customer"
	}

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	user.Password = "" // Clear password before returning
	return user, nil
}
