package admin

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
	store types.AdminStore
}

func NewHanlder(store types.AdminStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	// r.POST("/admins/signin", h.handleSignin)
	r.POST("/admins/signup", h.handleSignup)
}


func (h *Handler) handleSignup(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	body := types.AdminRegisterPayload{}

	data, err := ctx.GetRawData(); if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "User payload is not valid",
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

		user, _ := h.store.GetAdminByEmail(body.Email);

	if user.ID != "" {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"message": "Email already used",
		})
		return
	}

	hash, err := auth.HashPassword(body.Password); if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password",
		})
		return
	}

	err = h.store.CreateAdmin(types.Admin{
		ID: utils.GenerateUUID(),
		Username:   body.Username,
		Email:      body.Email,
		Password:   hash,
		CreatedAt: time.Now().Format(time.RFC3339),
	})

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Admin registered successfully",
	})
}