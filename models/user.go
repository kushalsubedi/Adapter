package models

// User represents a user entity in the system
type User struct {
	ID   int    `db:id, primary`
	Name string `db:"name"`
}
