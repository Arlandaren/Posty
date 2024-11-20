package config

import (
	"errors"
	"os"
)

type PostgresConfig struct {
	ConnStr string
}

func GetPostgres() (*PostgresConfig, error) {
	pgConn := os.Getenv("PG_STRING")
	if pgConn == "" {
		return nil, errors.New("not found PG_STRING")
	}
	return &PostgresConfig{
		ConnStr: pgConn,
	}, nil
}
