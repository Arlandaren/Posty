package repository

import (
	"log"
	"service/internal/shared/storage/postgres"
)

type Repository struct {
	db *postgres.Wrapper
}

func NewRepository(wrapper *postgres.Wrapper) *Repository {
	log.Println("NewRepository")
	return &Repository{
		db: wrapper,
	}
}
