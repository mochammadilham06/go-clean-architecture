package database

import (
	"database/sql"
	"fmt"
	"go-portfolio/server/lib/environment"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func ProvideSQLDatabase(config *environment.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME,
		"disable",      // sslmode, can be "disable", "require", etc.
		"Asia/Jakarta", // TimeZone, adjust as needed
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("[DB_INIT] : Failed set Connection: %v", err)
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	log.Printf("[DB_INIT] : Try to connect db: %s:%d...", config.DB_HOST, config.DB_PORT)
	if err := db.Ping(); err != nil {
		db.Close()
		log.Printf("[DB_INIT] (Ping): Failed to connect DB. Error: %v", err)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db.SetMaxIdleConns(config.DB_MAX_IDLE)
	db.SetMaxOpenConns(config.DB_MAX_IDLE)
	db.SetConnMaxLifetime(5 * time.Minute)

	// success connected to db
	log.Printf("DB Connected (Host: %s:%d). Pool: Idle=%d, Open=%d",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_MAX_IDLE,
		config.DB_MAX_OPEN,
	)

	return db, nil
}
