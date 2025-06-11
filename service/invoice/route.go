package invoice

import (
	"github.com/gin-gonic/gin"
	"github.com/rayhan889/lumbaumbah-backend/service/auth"
	"github.com/rayhan889/lumbaumbah-backend/types"
)

type Handler struct {
	store types.InvoiceStore
}

func NewHanlder(store types.InvoiceStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.Use(auth.Authenticate())
	r.GET("/payment/:id", h.handleGetPaymentByID)
}

func (h *Handler) handleGetPaymentByID(ctx *gin.Context) {

}
