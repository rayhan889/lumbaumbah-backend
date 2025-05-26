package main

import (
	"database/sql"
	"log"

	"github.com/rayhan889/lumbaumbah-backend/cmd/api"
	"github.com/rayhan889/lumbaumbah-backend/config"
	"github.com/rayhan889/lumbaumbah-backend/db"
)

func main() {
	dbConfig := config.Config{
		Host:     config.Envs.Host,
		Port:     config.Envs.Port,
		User:     config.Envs.User,
		Password: config.Envs.Password,
		DBName:   config.Envs.DBName,
		SSLMode:  config.Envs.SSLMode,
	}
	
	db, err := db.NewPostgresStorage(&dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	initDB(db)

	server := api.NewAPIServer(":8080", nil)
	err = server.Run(); if err != nil {
		panic(err)
	}
}

func initDB(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Connected to the database successfully")
}