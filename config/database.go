package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DatabaseConfig holds database connection parameters
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresConnection creates a new PostgreSQL database connection
func NewPostgresConnection(cfg DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// NewMySQLConnection creates a new MySQL database connection
func NewMySQLConnection(cfg DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
