package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rayhan889/lumbaumbah-backend/config"
	"github.com/rayhan889/lumbaumbah-backend/db"
)

func main() {
	dbConfig := config.Config{
		Host:     config.Envs.Host,
		DBPort:   config.Envs.DBPort,
		User:     config.Envs.User,
		Password: config.Envs.Password,
		DBName:   config.Envs.DBName,
		SSLMode:  config.Envs.SSLMode,
	}
	db, err := db.NewPostgresStorage(&dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	px, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}


	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		px,
	)
	if err != nil {
		log.Fatal(err)
	}

	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}