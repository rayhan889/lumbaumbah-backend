package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayhan889/lumbaumbah-backend/types"
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
	r.POST("/signin", h.handleSignin)
	r.POST("/signup", h.handleSignup)
}

func (h *Handler) handleSignin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
        "message": "user signin...",
    })
}

func (h *Handler) handleSignup(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
        "message": "user signup...",
    })
}