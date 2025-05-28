package address

import (
	"github.com/gin-gonic/gin"
	"github.com/rayhan889/lumbaumbah-backend/types"
)

type Handler struct {
	store types.AddressStore
}

func NewHanlder(store types.AddressStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/addresses", h.handleGetAddressByUserID)
	r.POST("/addresses/create", h.handleCreateAddress)
}

func (h *Handler) handleCreateAddress(ctx *gin.Context) {}

func (h *Handler) handleGetAddressByUserID(ctx *gin.Context) {}