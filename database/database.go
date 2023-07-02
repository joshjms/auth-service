package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joshjms/auth-service/models"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "auth"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Failed to connect to PostgreSQL database!")
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL database!")

	return db, nil
}

var DB *gorm.DB

func init() {
	var err error
	DB, err = ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB.AutoMigrate(&models.User{})
}
