package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Muhammad5943/notes-api-fiber/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declear variable for database
var DB *gorm.DB

// ConnectDB to connect DB
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println("Database Not Connected")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect to database")
	}

}
