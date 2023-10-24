package content

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/postgres"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/sqlxext"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/entity"
)

type Repository[T entity.Content] struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository[entity.Content] {
	r := new(Repository[entity.Content])
	r.db = db
	return r
}

func (r *Repository[T]) Create(e entity.Content, ctx context.Context) error {
	var lastId string
	q := postgres.BuildInsertQuery(tableName, []string{"name", "created_at", "updated_at"}, "RETURNING id")
	err := r.db.QueryRowContext(ctx, q, e.Name, e.CreatedAt, e.UpdatedAt).Scan(&lastId)
	if err != nil {
		return err
	}
	e.ID = lastId
	return nil
}

func (r *Repository[T]) ReadMany(limit, offset int, ctx context.Context) ([]entity.Content, error) {
	d := []entity.Content{}
	q := postgres.BuildSelectQuery(tableName, []string{}, []string{"is_deleted"}, "LIMIT $2 OFFSET $3")
	err := r.db.SelectContext(ctx, &d, q, limit, offset)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *Repository[T]) ReadOne(id string, ctx context.Context) (entity.Content, error) {
	b := entity.Content{}
	q := postgres.BuildSelectQuery(tableName, []string{}, []string{"id"}, "LIMIT $2")
	err := r.db.Get(&b, q, id)
	return b, err
}

func (r *Repository[T]) Update(id string, e entity.Content, ctx context.Context) (int64, error) {
	q := postgres.BuildUpdateQuery(tableName, []string{"name", "updated_at"}, []string{"id"}, "")
	res, err := r.db.Exec(q, id, e.Name, e.UpdatedAt)
	if err != nil {
		return -1, err
	}
	return sqlxext.GetRowsAffected(res), nil
}

func (r *Repository[T]) Delete(id string, ctx context.Context) (int64, error) {
	q := postgres.BuildDeleteQuery(tableName, []string{"id"}, "")
	res, err := r.db.Exec(q, id)
	if err != nil {
		return -1, err
	}
	return sqlxext.GetRowsAffected(res), nil
}

func (r *Repository[T]) DB() *sqlx.DB {
	return r.db
}

func (r *Repository[T]) createManyPQ(entities []entity.Content, ctx context.Context) error {
	txn, err := r.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := txn.Prepare(pq.CopyIn(tableName, "name", "created_at", "updated_at"))
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
