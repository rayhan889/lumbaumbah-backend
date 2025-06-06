package utils

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rayhan889/lumbaumbah-backend/config"
	"github.com/rayhan889/lumbaumbah-backend/types"
)

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

var Validate = validator.New()

func GetToken(ctx *gin.Context) string {
	const prefix = "Bearer "
	authHeader := ctx.GetHeader("Authorization")

	if !strings.HasPrefix(authHeader, prefix) {
		return ""
	}

	return strings.TrimSpace(authHeader[len(prefix):])
}

func VerifyToken(tokenString string) (claims *types.JWTClaims, msg string) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&types.JWTClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Envs.JWTSecret), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*types.JWTClaims)
	if !ok {
		msg = fmt.Sprintf("Invalid token claims")
		return
	}

	expTime, err := claims.GetExpirationTime()
	if err != nil {
		msg = err.Error()
		return
	}
	if expTime.Before(time.Now()) {
		msg = fmt.Sprintf("Token expired")
		return
	}

	return claims, msg
}

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetString("role")
		log.Printf("Role: %s", role)
		log.Printf("Allowed roles: %s", allowedRoles)
		if role == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "Missing role",
			})
			return
		}

		for _, allowed := range allowedRoles {
			if role == allowed {
				ctx.Next()
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Not allowed to access",
		})
	}
}