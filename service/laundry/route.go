package laundry

import (
	"encoding/json"
	"log"
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
	r.POST("/laundry/requests/create", utils.RequireRole("user"), h.handleCreateLaundryRequest)
	r.GET("/laundry/requests", utils.RequireRole("user"), h.handleGetLaundryRequestsByUserID)
}

func (h *Handler) handleGetLaundryRequestsByUserID(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	userId := ctx.GetString("user_id")
	
	requsts, err := h.store.GetLaundryRequestsByUseID(userId); if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get laundry requests",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"requests": requsts,
	})
}

func (h *Handler) handleGetLaundryTypes(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
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

func (h *Handler) handleCreateLaundryRequest(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	body := types.LaundryRequestPayload{}
	userId := ctx.GetString("user_id")

	data, err := ctx.GetRawData(); if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Request type payload is not valid",
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

	laundryType, err := h.store.GetLaundryTypeByID(body.LaundryTypeID); if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get laundry type",
		})
		return
	}

	completionDate := calculateCompletionDate(laundryType.EstimatedDays)
	var adminID *string = nil

	log.Printf("completionDate: %s", completionDate)

	err = h.store.CreateLaundryRequest(types.LaundryRequest{
		ID: utils.GenerateUUID(),
		UserID: userId,
		AdminID: adminID,
		LaundryTypeID: body.LaundryTypeID,
		AddressID: body.AddressID,
		Weight: body.Weight,
		Notes: body.Notes,
		Status: string(types.StatusPending),
		CompletionDate: completionDate,
	}); 
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create laundry request",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Laundry request created successfully",
	})
}