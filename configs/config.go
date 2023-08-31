package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var (
	conn *pgx.Conn
)

const (
	defaultPostgresUser     = "postgres"
	defaultPostgresPassword = "12345"
	defaultPostgresHost     = "postgres"
	defaultPostgresDB       = "dynamic-segmentation"
	defaultPostgresPort     = "5432"
)

func init() {
	initConnection()
}

func initConnection() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	user := getEnvOrFallback("POSTGRES_USER", defaultPostgresUser)
	password := getEnvOrFallback("POSTGRES_PASSWORD", defaultPostgresPassword)
	host := getEnvOrFallback("POSTGRES_HOST", defaultPostgresHost)
	db := getEnvOrFallback("POSTGRES_DB", defaultPostgresDB)
	port := getEnvOrFallback("POSTGRES_PORT", defaultPostgresPort)
	dburl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, db)
	conn, err = pgx.Connect(context.Background(), dburl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err)
	}
}

func Config() *pgx.Conn {
	if conn == nil {
		panic("unexpected conn == nil")
	}
	return conn
}

func getEnvOrFallback(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}
