package repository

import (
	"database/sql"
	"fmt"
	"project/models"
)

// MySQLRepo implements UserRepository for MySQL
type MySQLRepo struct {
	db *sql.DB
}

// NewMySQLRepo creates a new MySQL repository
func NewMySQLRepo(db *sql.DB) *MySQLRepo {
	return &MySQLRepo{db: db}
}

// Create inserts a new user into MySQL database
func (m *MySQLRepo) Create(user models.User) error {
	_, err := m.db.Exec(
		"INSERT INTO users (name) VALUES (?)",
		user.Name,
	)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

// GetAll retrieves all users from MySQL database
func (m *MySQLRepo) GetAll() ([]models.User, error) {
	rows, err := m.db.Query("SELECT id, name FROM users")
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
