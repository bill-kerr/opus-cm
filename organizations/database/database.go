package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the variable that holds the database instance.
var DB *gorm.DB

// Init returns an initialized database.
func Init(connString string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to Postgres database.")
	}
	fmt.Println("Organizations Service connected to PostgreSQL database.")
	DB = db
	return DB
}

// Migrate runs database migrations for all provided models.
func Migrate(db *gorm.DB, models ...interface{}) {
	if err := db.AutoMigrate(models); err != nil {
		log.Fatalln("Database migration error.")
	}
}

// GetDB returns a reference to the initialized database instance.
func GetDB() *gorm.DB {
	return DB
}