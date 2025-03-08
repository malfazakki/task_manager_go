package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// dsn Stands for "Data Source Name," a connection string used to connect to the PostgreSQL database.
	dsn := fmt.Sprintf("host=%s, user=%s, password=%s, dbname=%s, port=%s, sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	/*
		- gorm.Open(...): Opens a connection to the database using the PostgreSQL driver.
		- postgres.Open(dsn): Passes the DSN string to GORM’s PostgreSQL driver.
		- &gorm.Config{}: Initializes GORM with default configurations.
		- var err error: Declares an err variable to capture any errors that might occur during database connection.
	*/
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to Database: ", err)
	}
}
