package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rayhan889/lumbaumbah-backend/config"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

func GenerateJWT(userID string, secret []byte, role string) (string, error) {
	expTime := time.Second * time.Duration(config.Envs.JWTExpirationsInSecond)

	token :=jwt.NewWithClaims(JWT_SIGNING_METHOD, jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(expTime).Unix(),
		"role": role,
	})

	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}