package main

import (
	"github.com/rs/cors"
)

func allowCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // ваш фронтенд адрес
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
}
