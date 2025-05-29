package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rayhan889/lumbaumbah-backend/config"
	"github.com/rayhan889/lumbaumbah-backend/types"
	"github.com/rayhan889/lumbaumbah-backend/utils"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

func GenerateJWT(userID string, secret []byte, role string) (string, error) {
	expTime := time.Second * time.Duration(config.Envs.JWTExpirationsInSecond)
	claims := &types.JWTClaims{
		UserID: userID,
		Role:   role,
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Add(expTime).Unix(),
		},
	}

	token :=jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)

	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := utils.GetToken(ctx)
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Missing token",
			})
			return
		}

		claims, msg := utils.VerifyToken(token)
		if msg != "" {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": msg,
			})
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}