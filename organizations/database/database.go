package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database represents the database instance.
type Database struct {
	*gorm.DB
}

// DB is the variable that holds the database instance.
var DB *gorm.DB


// Init returns an initialized database.
func Init(connString string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connected to Postgres database.")
	}
	fmt.Println("Organizations Service connected to PostgreSQL database.")
	DB = db
	return DB
}

// Migrate runs database migrations for all models.
func Migrate(db *gorm.DB) {
	return
}