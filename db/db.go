package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rayhan889/lumbaumbah-backend/config"
)

func NewPostgresStorage(cfg *config.Config) (*sql.DB, error) {
	connString := cfg.FormatDSN()
	db, err := sql.Open("pgx/v5", connString)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	return db, nil
}