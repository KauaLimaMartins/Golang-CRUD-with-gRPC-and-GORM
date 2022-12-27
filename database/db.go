package database

import (
	"fmt"
	"log"
	"os"

	"github.com/kaualimamartins/gRPC-Go-Gorm/customers/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectWithDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", dbHost, dbUser, dbPass, dbName)
	// dsn := "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable port=5432"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("failed to open connection to database. ", err.Error())
	}

	err = DB.AutoMigrate(&models.Customer{})

	if err != nil {
		log.Panic("failed to migrate database", err.Error())
	}

	log.Print("Successfully connected to database")
}
