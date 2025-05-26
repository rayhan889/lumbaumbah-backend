package db

import (
	"log"

	"github.com/rayhan889/lumbaumbah-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresStorage(cfg *config.Config) (*gorm.DB, error) {
	connString := cfg.FormatDSN()
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	return db, nil
}