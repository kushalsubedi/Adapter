package repository

import (
	"database/sql"
	"fmt"

	"project/models"
)

// PostgresRepo implements UserRepository for PostgreSQL
type PostgresRepo struct {
	db *sql.DB
}

// NewPostgresRepo creates a new PostgreSQL repository
func NewPostgresRepo(db *sql.DB) (*PostgresRepo, error) {
	repo := &PostgresRepo{db: db}

	// auto-migrate on startup
	if err := repo.AutoMigrate(models.User{}); err != nil {
		return nil, err
	}

	return repo, nil
}

// Create inserts a new user into PostgreSQL database
func (p *PostgresRepo) Create(user models.User) error {
	res, err := p.db.Exec(
		"INSERT INTO users (name) VALUES ($1)",
		user.Name,
	)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	fmt.Println("Inserted rows:", rows)
	return nil
}

// GetAll retrieves all users from PostgreSQL database
func (p *PostgresRepo) GetAll() ([]models.User, error) {
	rows, err := p.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return users, nil
}
