package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte(os.Getenv("SECRET"))

// Token generate karna
func GenerateJWT(userID uint, role string) (string, error) {
    claims := jwt.MapClaims{}
    claims["user_id"] = userID
    claims["role"] = role
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24h expiry

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(JWTSecret)
}
