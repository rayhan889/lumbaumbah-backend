package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rayhan889/lumbaumbah-backend/config"
	"github.com/rayhan889/lumbaumbah-backend/service/auth"
	"github.com/rayhan889/lumbaumbah-backend/types"
	"github.com/rayhan889/lumbaumbah-backend/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHanlder(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/users/signin", h.handleSignin)
	r.POST("/users/signup", h.handleSignup)
}

func (h *Handler) handleSignin(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	body := types.SigninPayload{}

	data, err := ctx.GetRawData(); if err != nil || len(data) == 0 {
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

	user, err := h.store.GetUserByEmail(body.Email); if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Failed to get user",
		})
		return
	}

	if user.ID == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	if !auth.CheckPassword(body.Password, user.Password) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.GenerateJWT(user.ID, secret, "user"); if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate JWT token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
        "token": token,
    })
}

func (h *Handler) handleSignup(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	body := types.UserRegisterPayload{}

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

	user, _ := h.store.GetUserByEmail(body.Email);

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

	err = h.store.CreateUser(types.User{
		ID: utils.GenerateUUID(),
		Username:   body.Username,
		FullName:  body.FullName,
		Email:      body.Email,
		Password:   hash,
		PhoneNumber: body.PhoneNumber,
		CreatedAt: time.Now().Format(time.RFC3339),
	})

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}