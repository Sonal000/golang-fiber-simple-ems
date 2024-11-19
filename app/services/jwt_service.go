package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(userID string, role string) (string, error)
	ValidateToken(token string) (map[string]interface{}, error)
}

type jwtService struct {
	secretKey string
}

func NewJWTService() JWTService {

	secretKey := os.Getenv("JWT_SECRET")
	return &jwtService{
		secretKey: secretKey,
	}
}

func (j *jwtService) GenerateToken(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("unable to parse claims")
	}

	return claims, nil
}
