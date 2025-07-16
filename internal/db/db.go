package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Init() error {
	// Load DB_URL from config loaded from .env
	dsn := os.Getenv("DB_URL")
	fmt.Println("🔑 Inside db.Init() ➔ DB_URL:", dsn)  
	if dsn == "" {
		return fmt.Errorf("❌ Database URL is missing")
	}

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return fmt.Errorf("❌ Failed to parse DB config: %w", err)
	}

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("❌ Failed to create DB pool: %w", err)
	}

	// Check DB connectivity
	if err := DB.Ping(context.Background()); err != nil {
		return fmt.Errorf("❌ Database not reachable: %w", err)
	}

	log.Println("✅ Connected to PostgreSQL")
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
		log.Println("✅ Database connection closed")
	}
}
