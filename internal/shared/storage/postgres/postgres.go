package postgres

import (
	"context"
	"service/internal/shared/config"

	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func InitPostgres(cfg *config.PostgresConfig, maxRetries int) (*pgxpool.Pool, error) {
	var retryCount int

	pool, err := pgxpool.New(context.TODO(), cfg.ConnStr)
	for err != nil && retryCount < maxRetries {
		fmt.Println("retryCount", retryCount)
		time.Sleep(time.Second * 6)
		pool, err = pgxpool.New(context.TODO(), cfg.ConnStr)
		retryCount++
	}

	if err != nil {
		return nil, errors.New("Failed to connect to Postgres: " + err.Error())
	}

	retryCount = 0
	err = pool.Ping(context.TODO())

	for err != nil && retryCount < maxRetries {
		fmt.Println(err, "Try to connect")
		time.Sleep(time.Second * 5)
		err = pool.Ping(context.TODO())
		retryCount++
	}

	if err != nil {
		return nil, err
	}

	return pool, nil
}
