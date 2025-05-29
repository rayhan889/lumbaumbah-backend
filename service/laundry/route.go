package laundry

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
	store types.LaundryStore
}

func NewHanlder(store types.LaundryStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.Use(auth.Authenticate())
	r.POST("/laundry/types/create", utils.RequireRole("admin"), h.hanldeCreateLaundryType)
	r.GET("/laundry/types", h.handleGetLaundryTypes)
}


func (h *Handler) handleGetLaundryTypes(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	types, err := h.store.GetLaundryTypes(); if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get laundry types",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"types": types,
	})
}

func (h *Handler) hanldeCreateLaundryType(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	body := types.LaundryTypePayload{}

	data, err := ctx.GetRawData(); if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Laundry type payload is not valid",
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

	err = h.store.CreateLaundryType(types.LaundryType{
		ID: utils.GenerateUUID(),
		Name: body.Name,
		Description: body.Description,
		Price: body.Price,
		EstimatedDays: body.EstimatedDays,
		IsActive: true,
		CreatedAt: time.Now().Format(time.RFC3339),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create laundry type",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Laundry type created successfully",
	})
}