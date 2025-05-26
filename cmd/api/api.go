package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rayhan889/lumbaumbah-backend/service/user"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (s *APIServer) Run() error {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	userStore := user.NewStore(s.db)
	userService := user.NewHanlder(userStore)
	userService.RegisterRoutes(v1)

	return r.Run(s.address)
}