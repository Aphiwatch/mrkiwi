package database

import (
	"context"
	"database/sql"
	"fmt"

	// "log"

	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	database = os.Getenv("POSTGRES_DB")
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	host     = os.Getenv("POSTGRES_HOST")
	port     = os.Getenv("POSTGRES_PORT")
)

func ConnectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}

// func GetDB() *sql.DB {
// 	if db == nil {
// 		log.Fatal("Database is not initialized")
// 	}
// 	return db
// }
