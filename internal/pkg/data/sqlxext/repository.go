package sqlxext

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository[T any] interface {
	Create(e T, ctx context.Context) error

	ReadMany(limit, offset int, ctx context.Context) ([]T, error)

	ReadOne(id string, ctx context.Context) (T, error)

	Update(id string, e T, ctx context.Context) (int64, error)

	Delete(id string, ctx context.Context) (int64, error)

	DB() *sqlx.DB
}
