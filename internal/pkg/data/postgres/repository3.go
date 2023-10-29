package postgres

import (
	"context"
	"database/sql"
)

type Repository3[T any] interface {
	Create(e T, ctx context.Context) (string, error)

	ReadMany(limit, offset int, ctx context.Context) (*sql.Rows, error)

	ReadOne(id string, ctx context.Context) *sql.Row

	Update(id string, e T, ctx context.Context) (int64, error)

	Delete(id string, ctx context.Context) (int64, error)

	DB() *sql.DB
}
