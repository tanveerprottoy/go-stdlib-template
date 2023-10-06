package content

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/sqlxpkg"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/entity"
)

const TableName = "contents"

type Repository[T entity.Content] struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository[entity.Content] {
	r := new(Repository[entity.Content])
	r.db = db
	return r
}

func (r *Repository[T]) Create(e *entity.Content) error {
	var lastId string
	err := r.db.QueryRow("INSERT INTO "+TableName+" (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id", e.Name, e.CreatedAt, e.UpdatedAt).Scan(&lastId)
	if err != nil {
		return err
	}
	e.Id = lastId
	return nil
}

func (r *Repository[T]) ReadMany(limit, offset int) ([]entity.Content, error) {
	d := []entity.Content{}
	err := r.db.Select(&d, "SELECT * FROM "+TableName+" LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *Repository[T]) ReadOne(id string) (entity.Content, error) {
	b := entity.Content{}
	err := r.db.Get(&b, "SELECT * FROM "+TableName+" WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (r *Repository[T]) Update(id string, e *entity.Content) (int64, error) {
	q := "UPDATE " + TableName + " SET name = $2, updated_at = $3 WHERE id = $1"
	res, err := r.db.Exec(q, id, e.Name, e.UpdatedAt)
	if err != nil {
		return -1, err
	}
	return sqlxpkg.GetRowsAffected(res), nil
}

func (r *Repository[T]) Delete(id string) (int64, error) {
	q := "DELETE FROM " + TableName + " WHERE id = $1"
	res, err := r.db.Exec(q, id)
	if err != nil {
		return -1, err
	}
	return sqlxpkg.GetRowsAffected(res), nil
}

func (r *Repository[T]) createMany(entities []entity.Content, ctx context.Context) error {
	txn, err := r.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := txn.Prepare(pq.CopyIn(TableName, "name", "created_at", "updated_at"))
	if err != nil {
		return (err)
	}
	// close the statement when done
	defer stmt.Close()
	for _, e := range entities {
		_, err := stmt.Exec(e.Name, e.CreatedAt, e.UpdatedAt)
		if err != nil {
			txn.Rollback()
			return err
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	err = txn.Commit()
	if err != nil {
		return err
	}
	return nil
}
