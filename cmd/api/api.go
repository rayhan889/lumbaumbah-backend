package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rayhan889/lumbaumbah-backend/service/address"
	"github.com/rayhan889/lumbaumbah-backend/service/admin"
	"github.com/rayhan889/lumbaumbah-backend/service/invoice"
	"github.com/rayhan889/lumbaumbah-backend/service/laundry"
	"github.com/rayhan889/lumbaumbah-backend/service/user"
	"gorm.io/gorm"
)

type APIServer struct {
	address string
	db      *gorm.DB
}

func NewAPIServer(address string, db *gorm.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (s *APIServer) Run() error {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := r.Group("/api/v1")

	userStore := user.NewStore(s.db)
	userService := user.NewHanlder(userStore)
	userService.RegisterRoutes(v1)

	adminStore := admin.NewStore(s.db)
	adminService := admin.NewHanlder(adminStore)
	adminService.RegisterRoutes(v1)

	addressStore := address.NewStore(s.db)
	addressService := address.NewHanlder(addressStore)
	addressService.RegisterRoutes(v1)

	laundryStore := laundry.NewStore(s.db)
	laundryService := laundry.NewHanlder(laundryStore)
	laundryService.RegisterRoutes(v1)

	invoiceStore := invoice.NewStore(s.db)
	invoiceService := invoice.NewHanlder(invoiceStore)
	invoiceService.RegisterRoutes(v1)

	return r.Run(s.address)
}
