package statistic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayhan889/lumbaumbah-backend/service/auth"
	"github.com/rayhan889/lumbaumbah-backend/types"
	"github.com/rayhan889/lumbaumbah-backend/utils"
)

type Handler struct {
	store types.StatisticStore
}

func NewHanlder(store types.StatisticStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.Use(auth.Authenticate())
	r.GET("/statistics/user", utils.RequireRole("user"), h.handleGetUserStatistics)
	r.GET("/statistics/admin", utils.RequireRole("admin"), h.handleGetAllStatistics)
	r.GET("/users-lists", utils.RequireRole("admin"), h.handleGetUserLists)
}

func (h *Handler) handleGetUserLists(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	users, err := h.store.GetUsersList()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get user lists",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *Handler) handleGetUserStatistics(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	userId := ctx.GetString("user_id")

	statistics, err := h.store.GetUserStatistics(userId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"statistics": statistics,
	})
}

func (h *Handler) handleGetAllStatistics(ctx *gin.Context) {

}
