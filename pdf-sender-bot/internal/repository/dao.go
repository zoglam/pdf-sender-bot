package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DAO interface {
	NewUserQuery() UserQuery
}
type dao struct{}

var DB *pgxpool.Pool

func NewDAO() DAO {
	return &dao{}
}

func NewDB(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	var err error
	DB, err = pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return DB, nil
}

func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{}
}
