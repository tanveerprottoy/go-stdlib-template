package content

import (
	"database/sql"
	"fmt"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/postgres"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/entity"
)

type RepositorySQL[T entity.Content] struct {
	db *sql.DB
}

func NewRepositorySQL(db *sql.DB) *RepositorySQL[entity.Content] {
	return &RepositorySQL[entity.Content]{db: db}
}

func (r *RepositorySQL[T]) Create(e *entity.Content) error {
	res, err := r.db.Exec("INSERT INTO "+TableName+" (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id", e.Name, e.CreatedAt, e.UpdatedAt)
	if err != nil {
		return err
	}
	fmt.Println("res: ", res)
	return nil
}

func (r *RepositorySQL[T]) ReadMany(limit, offset int) (*sql.Rows, error) {
	rows, err := r.db.Query("SELECT * FROM "+TableName+" LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *RepositorySQL[T]) ReadOne(id string) *sql.Row {
	return r.db.QueryRow("SELECT * FROM "+TableName+" WHERE id = $1 LIMIT 1", id)
}

func (r *RepositorySQL[T]) Update(id string, e *entity.Content) (int64, error) {
	q := "UPDATE " + TableName + " SET name = $2, updated_at = $3 WHERE id = $1"
	res, err := r.db.Exec(q, id, e.Name, e.UpdatedAt)
	if err != nil {
		return -1, err
	}
	return postgres.GetRowsAffected(res), nil
}

func (r *RepositorySQL[T]) Delete(id string) (int64, error) {
	q := "DELETE FROM " + TableName + " WHERE id = $1"
	res, err := r.db.Exec(q, id)
	if err != nil {
		return -1, err
	}
	return postgres.GetRowsAffected(res), nil
}
