package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgxpool.Pool, error) {
	err := godotenv.Load() // load .env

	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Get the connection string from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	println(dbHost, dbPort, dbUser, dbPassword, dbName)

	connectionString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	// Connect to the database using pgxpool
	dbpool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %w", err)
	}

	fmt.Println("Successfully connected to the database!")
	return dbpool, nil
}
