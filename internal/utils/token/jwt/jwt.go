package jwt

import (
	"cmd/poker-backend/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userId uint) (string, error) {
	cfg := config.Get()

	claims := jwt.MapClaims{
		"iss": cfg.Auth.JwtToken.Issuer,
		"sub": userId,
		"exp": time.Now().Add(time.Hour * time.Duration(cfg.Auth.JwtToken.LifespanHours)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.Secrets.Application.ApplicationSecretKey))
}

func isTokenValid() {}

func GetUserIDFromToken() {}

func GetTokenFromHeader() {}
