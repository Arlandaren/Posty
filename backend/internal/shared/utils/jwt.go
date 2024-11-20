package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"service/internal/shared/config"
	"service/internal/shared/storage/dto"
)

func ValidateToken(tokenString string) (*dto.Claims, error) {
	secretKey := []byte(config.GetJwt())

	claims := &dto.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
