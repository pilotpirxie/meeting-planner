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

func Init(ctx context.Context, dbURL string) (*DB, error) {
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set; set it in environment or .env")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse database URL: %w", err)
	}
	config.MaxConns = 5
	config.MinConns = 5
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("Unable to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %w", err)
	}

	return &DB{
		Pool:    pool,
		Queries: sqlc.New(pool),
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
