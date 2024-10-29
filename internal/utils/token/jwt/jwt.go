package jwt

import (
	"cmd/poker-backend/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
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

type InvalidTokenHeader struct {
	err string
}

func (e *InvalidTokenHeader) Error() string {
	return fmt.Sprintf(e.err)
}

func getTokenStringFromHeader(c *gin.Context) (string, error) {
	header := c.GetHeader("Authorization")

	parts := strings.Split(header, " ")

	if len(parts) != 2 {
		return "", &InvalidTokenHeader{fmt.Sprintf("Expected '<Type> <token>' format. Got %s", header)}
	}

	if !strings.Contains(parts[0], "Bearer") {
		return "", &InvalidTokenHeader{fmt.Sprintf("Expected 'Bearer' token. Got %s", header)}
	}

	return parts[1], nil
}

func GetToken(c *gin.Context) (*jwt.Token, error) {
	tokenString, err := getTokenStringFromHeader(c)

	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, tokenValidationFunc)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token [%s] is invalid", tokenString)
	}

	return token, nil
}

func tokenValidationFunc(token *jwt.Token) (interface{}, error) {
	cfg := config.Get()

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(cfg.Secrets.Application.ApplicationSecretKey), nil
}

func GetUserIDFromToken(token *jwt.Token) (int, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, nil
	}

	return int(claims["sub"].(float64)), nil
}
