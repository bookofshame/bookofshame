package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/bookofshame/bookofshame/pkg/config"
	"github.com/bookofshame/bookofshame/pkg/logging"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Sql struct {
	*sqlx.DB
}

func New(ctx context.Context, cfg config.Config) (*Sql, error) {
	logger := logging.FromContext(ctx)

	if cfg.TursoDbUrl == "" {
		return nil, fmt.Errorf("database host url missing")
	}

	if cfg.Env != "development" && cfg.TursoDbAuthToken == "" {
		return nil, fmt.Errorf("database auth token missing")
	}

	connectionString := cfg.TursoDbUrl
	if cfg.TursoDbAuthToken != "" {
		connectionString = fmt.Sprintf("%s?authToken=%s", cfg.TursoDbUrl, cfg.TursoDbAuthToken)
	}

	logger.Debugln(connectionString)

	db, err := sqlx.Connect("libsql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database. error: %w", err)
	}

	// check if connection is usable
	var placeholder int
	if err := db.Get(&placeholder, "SELECT 1"); err != nil {
		return nil, err
	}

	logger.Debugln("connected to database")

	return &Sql{db}, nil
}
