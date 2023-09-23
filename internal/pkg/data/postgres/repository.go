package postgres

import (
	"context"
	"database/sql"
)

type Repository[T any] interface {
	Create(e *T) (string, error)

	ReadMany(limit, offset int) (*sql.Rows, error)

	ReadOne(id string) *sql.Row

	Update(id string, e *T) (int64, error)

	Delete(id string, ctx context.Context) (int64, error)

	DB() *sql.DB

	// TableName() string
}
