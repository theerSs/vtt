package database

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/theerSs/vtt/internal/platform/env"
)

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(getConnectionURL())
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = 30 * time.Minute
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.HealthCheckPeriod = 30 * time.Second

	return pgxpool.NewWithConfig(ctx, cfg)
}

func getConnectionURL() string {
	user := url.QueryEscape(env.PostgresUser.GetValue())
	pass := url.QueryEscape(env.PostgresPassword.GetValue())

	host := env.PostgresHost.GetValue()
	port := env.PostgresPort.GetValue()
	db := env.PostgresDb.GetValue()

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		user, pass, host, port, db,
	)
}
