package database

import (
	"cmd/poker-backend/internal/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(cfg config.DatabaseConfig) *gorm.DB {
	dsn := cfg.MakeDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("[!] Can't connect to the database")
		log.Fatalf("[!] %e", err)
	}

	return db
}
