package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func PostgreSql(ctx context.Context, logger *zap.Logger) *pgxpool.Pool {
	logger.Info("initializing postgresql connection")
	cfg, err := NewPostgresConfig()
	if err != nil {
		logger.Error("error while initializing database connection", zap.Error(err))
		panic(err)
	}

	ss := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?search_path=%s&pool_max_conns=40&pool_min_conns=1&pool_max_conn_lifetime=10s&pool_max_conn_idle_time=5s",
		cfg.GetUsername(),
		cfg.GetPassword(),
		cfg.GetHost(),
		cfg.GetPort(),
		cfg.GetDatabase(),
		cfg.GetSchema(),
	)

	dbase, err := pgxpool.New(ctx, ss)
	if err != nil {
		logger.Error("error when try to open database", zap.Error(err))
		panic(err)
	}

	return dbase
}
