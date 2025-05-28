package main

import (
	"fmt"
	"log"

	"github.com/rayhan889/lumbaumbah-backend/cmd/api"
	"github.com/rayhan889/lumbaumbah-backend/config"
	"github.com/rayhan889/lumbaumbah-backend/db"
	"gorm.io/gorm"
)

func main() {
	dbConfig := config.Config{
		Host:     config.Envs.Host,
		DBPort:     config.Envs.DBPort,
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

	port :=fmt.Sprintf(":%s", config.Envs.Port)
	server := api.NewAPIServer(port, db)
	err = server.Run(); if err != nil {
		panic(err)
	}
}

func initDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	err = sqlDB.Ping()

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Connected to the database successfully")
}