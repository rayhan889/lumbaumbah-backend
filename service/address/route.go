package address

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	r.Use(auth.Authenticate())
	r.GET("/addresses", utils.RequireRole("user"), h.handleGetAddressByUserID)
	r.POST("/addresses/create", utils.RequireRole("user"), h.handleCreateAddress)
}

func (h *Handler) handleCreateAddress(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	body := types.UserAddressPayload{}

	data, err := ctx.GetRawData(); if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Address payload is not valid",
		})
		return
	}

	err = json.Unmarshal(data, &body); if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request payload",
		})
		return
	}

	err = utils.Validate.Struct(body); if err != nil {
		errors := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
			"errors": errors.Error(),
		})
		return
	}

	userId := ctx.GetString("user_id")

	err = h.store.CreateAddress(types.Address{
		ID: utils.GenerateUUID(),
		UserID: userId,
		StreetAddress: body.StreedAddress,
		City: body.City,
		State: body.State,
		IsDefault: false,
		CreatedAt: time.Now().Format(time.RFC3339),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create address",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Address created successfully",
	})
}

func (h *Handler) handleGetAddressByUserID(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	userId := ctx.GetString("user_id")

	addresses, err := h.store.GetAddressesByUserID(userId); if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get addresses",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"addresses": addresses,
	})
}