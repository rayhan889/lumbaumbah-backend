package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayhan889/lumbaumbah-backend/service/auth"
	"github.com/rayhan889/lumbaumbah-backend/types"
	"github.com/rayhan889/lumbaumbah-backend/utils"
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
	r.Use(
		auth.Authenticate(),
		utils.RequireRole("user"),
	)
	{
		r.GET("/addresses", h.handleGetAddressByUserID)
		r.POST("/addresses/create", h.handleCreateAddress)
	}
}

func (h *Handler) handleCreateAddress(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Address created successfully",
	})
}

func (h *Handler) handleGetAddressByUserID(ctx *gin.Context) {
	userId := ctx.GetString("user_id")
	role := ctx.GetString("role")

	if userId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "User ID is required",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"userId": userId,
		"role": role,
	})
}