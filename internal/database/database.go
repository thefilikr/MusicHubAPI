package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"test-task-filikr/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

func SetupDB(config config.ConfigDB, log *slog.Logger) (*sql.DB, error) {

	connStr := fmt.Sprintf(
		"%s:%s@%s:%d/%s?sslmode=disable",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.NameDB,
	)

	db, err := sql.Open("postgres", "postgres://"+connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping database: %v", err)
	}

	log.Debug("Connect to DB")

	m, err := migrate.New(
		"file://migrations",
		"postgres://"+connStr)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize migrate: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("Failed to apply migrations: %v", err)
	}

	log.Info("Migrations applied successfully!")

	return db, nil
}
