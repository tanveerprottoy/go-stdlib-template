package postgres

import (
	"context"
	"database/sql"
)

type Repository1[T any] interface {
	Create(ctx context.Context, e T, args ...any) (string, error)

	ReadMany(ctx context.Context, limit, offset int, args ...any) (*sql.Rows, error)

	ReadOne(ctx context.Context, id string, args ...any) *sql.Row

	Update(ctx context.Context, id string, e T, args ...any) (int64, error)

	Delete(ctx context.Context, id string, args ...any) (int64, error)

	DB() *sql.DB
}
