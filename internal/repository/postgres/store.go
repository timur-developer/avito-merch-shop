package postgres

import (
	"avito-merch-shop/config"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QueryExecutor interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Ping(ctx context.Context) error
	Close()
}

type Store struct {
	pool QueryExecutor
	sb   squirrel.StatementBuilderType
}

func NewStore(cfg *config.Config) (*Store, error) {
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &Store{
		pool: pool,
		sb:   squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

func (s *Store) Pool() QueryExecutor {
	return s.pool
}

func (s *Store) Builder() squirrel.StatementBuilderType {
	return s.sb
}

func (s *Store) Close() {
	s.pool.Close()
}
