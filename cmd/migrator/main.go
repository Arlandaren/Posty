package main

import (
	"service/internal/shared/config"

	"errors"

	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {
	var migrationsPath, migrationsTable string

	cfg, err := config.GetPostgres()
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	pgString := cfg.ConnStr

	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse()

	if migrationsPath == "" {
		log.Fatal("migrations-path is required")
	}

	pgStringWithTable := fmt.Sprintf("%s&x-migrations-table=%s", pgString, migrationsTable)

	m, err := migrate.New(
		"file://"+migrationsPath,
		pgStringWithTable,
	)
	if err != nil {
		log.Fatalf("failed to initialize migrate instance: %v", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}
		log.Fatalf("failed to apply migrations: %v", err)
	}

	fmt.Println("migrations applied successfully")
}

type Log struct {
	verbose bool
}

func (l *Log) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (l *Log) Verbose() bool {
	return l.verbose
}
