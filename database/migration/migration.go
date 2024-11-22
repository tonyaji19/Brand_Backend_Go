package database

import (
	"log"

	"golang-api-service/config"

	"github.com/pressly/goose/v3"
)

func RunMigrations() {
	db, err := config.GetDBConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB: %v", err)
	}

	if err := goose.SetDialect("mssql"); err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	if err := goose.Up(sqlDB, "./database/migration"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations applied successfully")
}
