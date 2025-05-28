package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayhan889/lumbaumbah-backend/service/auth"
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
	r.Use(auth.Authenticate())
	{
		r.GET("/addresses", h.handleGetAddressByUserID)
		r.POST("/addresses/create", h.handleCreateAddress)
	}
}

func (h *Handler) handleCreateAddress(ctx *gin.Context) {
	userId := ctx.Query("user_id")

	if userId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "User ID is required",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"userId": userId,
	})
}

func (h *Handler) handleGetAddressByUserID(ctx *gin.Context) {
	userId := ctx.GetString("user_id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": userId,
	})
}