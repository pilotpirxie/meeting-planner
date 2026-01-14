package db

import (
	"context"
	"fmt"
	"meeting-planner/backend/internal/db/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool    *pgxpool.Pool
	Queries *sqlc.Queries
}

func Init(ctx context.Context, databaseURL string) (*DB, error) {
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set; set it in environment or .env")
	}

	poolConfig, configParsingError := pgxpool.ParseConfig(databaseURL)
	if configParsingError != nil {
		return nil, fmt.Errorf("Unable to parse database URL: %w", configParsingError)
	}
	poolConfig.MaxConns = 5
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = 30 * time.Minute
	poolConfig.MaxConnIdleTime = 5 * time.Minute
	poolConfig.HealthCheckPeriod = 1 * time.Minute

	connectionPool, poolCreationError := pgxpool.NewWithConfig(ctx, poolConfig)
	if poolCreationError != nil {
		return nil, fmt.Errorf("Unable to create connection pool: %w", poolCreationError)
	}

	if pingError := connectionPool.Ping(ctx); pingError != nil {
		return nil, fmt.Errorf("Unable to connect to database: %w", pingError)
	}

	return &DB{
		Pool:    connectionPool,
		Queries: sqlc.New(connectionPool),
	}, nil
}

func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}

func (db *DB) Healthcheck(ctx context.Context) error {
	return db.Pool.Ping(ctx)
}
