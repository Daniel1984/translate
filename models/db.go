package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, user, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err.Error())
	}

	db = conn
	// db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration
}

// GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
