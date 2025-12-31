package main

import (
	"fmt"
	"log"

	"project/config"
	"project/repository"
	"project/service"
)

func main() {
	// Configure database connection
	dbConfig := config.DatabaseConfig{
		Host:     "localhost",
		Port:     5433,
		User:     "postgres",
		Password: "postgres",
		DBName:   "appdb",
		SSLMode:  "disable",
	}

	// Create database connection
	db, err := config.NewPostgresConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repository
	repo := repository.NewPostgresRepo(db)
	
	// Uncomment to use MySQL instead:
	// mysqlConfig := config.DatabaseConfig{
	// 	Host:     "localhost",
	// 	Port:     3306,
	// 	User:     "root",
	// 	Password: "password",
	// 	DBName:   "appdb",
	// }
	// mysqlDB, err := config.NewMySQLConnection(mysqlConfig)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to MySQL: %v", err)
	// }
	// defer mysqlDB.Close()
	// repo = repository.NewMySQLRepo(mysqlDB)

	// Initialize service
	userService := service.NewUserService(repo)

	// Register users (uncomment to use)
	// if err := userService.RegisterUser("Kushal"); err != nil {
	// 	log.Printf("Failed to register user: %v", err)
	// }
	// if err := userService.RegisterUser("DevOps"); err != nil {
	// 	log.Printf("Failed to register user: %v", err)
	// }

	// List all users
	users, err := userService.ListUsers()
	if err != nil {
		log.Fatalf("Failed to list users: %v", err)
	}

	fmt.Println("Registered Users:", users)
}
