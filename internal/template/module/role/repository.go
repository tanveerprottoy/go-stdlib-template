package role

import (
	"context"
	"database/sql"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/postgres"
	"github.com/tanveerprottoy/stdlib-go-template/internal/workersinsights/module/role/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/timeext"
)

const tableName = "roles"

type Repository[T entity.Role] struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository[entity.Role] {
	return &Repository[entity.Role]{db: db}
}

func (r *Repository[T]) Create(e *entity.Role) (string, error) {
	var lastID string
	err := r.db.QueryRow("INSERT INTO "+tableName+" (name, key, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id", e.Name, e.Key, e.CreatedAt, e.UpdatedAt).Scan(&lastID)
	if err != nil {
		return lastID, err
	}
	return lastID, nil
}

func (r *Repository[T]) ReadMany(limit, offset int) (*sql.Rows, error) {
	rows, err := r.db.Query("SELECT * FROM "+tableName+" LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *Repository[T]) ReadOne(id string) *sql.Row {
	return r.db.QueryRow("SELECT * FROM "+tableName+" WHERE id = $1 LIMIT 1", id)
}

func (r *Repository[T]) Update(id string, e *entity.Role) (int64, error) {
	q := "UPDATE " + tableName + " SET name = $2, key = $3, updated_at = $4 WHERE id = $1"
	res, err := r.db.Exec(q, id, e.Name, e.Key, e.UpdatedAt)
	if err != nil {
		return -1, err
	}
	return postgres.GetRowsAffected(res), nil
}

func (r *Repository[T]) Delete(id string, ctx context.Context) (int64, error) {
	q := "UPDATE " + tableName + " SET is_deleted = $2, updated_at = $3 WHERE id = $1"
	res, err := r.db.Exec(q, id, true, timeext.NowUnix())
	if err != nil {
		return -1, err
	}
	return postgres.GetRowsAffected(res), nil
}

func (r *Repository[T]) DB() *sql.DB {
	return r.db
}
