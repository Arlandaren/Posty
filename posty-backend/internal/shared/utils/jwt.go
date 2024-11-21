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

func ParseTokenWithoutVerification(tokenStr string) (*dto.Claims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, &dto.Claims{})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*dto.Claims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}
