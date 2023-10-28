package postgres

import (
	"context"
	"database/sql"
)

// Repository2 if implemented should perform data conversion
// from sql type to desired struct type
type Repository2[T any] interface {
	Create(ctx context.Context, e T, args ...any) (string, error)

	ReadMany(ctx context.Context, limit, offset int, args ...any) ([]T, error)

	ReadOne(ctx context.Context, id string, args ...any) (T, error)

	Update(ctx context.Context, id string, e T, args ...any) (int64, error)

	Delete(ctx context.Context, id string, args ...any) (int64, error)

	DB() *sql.DB
}
