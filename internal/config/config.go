package config

import (
	"errors"
	"os"
)

var JWT_SECRET string
var DB_URL string

func LoadConfig() error {
	JWT_SECRET = os.Getenv("JWT_SECRET")
	DB_URL = os.Getenv("DB_URL")

	if JWT_SECRET == "" || DB_URL == "" {
		return errors.New("missing JWT_SECRET or DB_URL in environment")
	}
	return nil
}
