package content

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	pgxstdlib "github.com/jackc/pgx/v5/stdlib"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/postgres"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/entity"
)

const tableName = "contents"

type RepositorySQL[T entity.Content] struct {
	db *sql.DB
}

func NewRepositorySQL(db *sql.DB) *RepositorySQL[entity.Content] {
	return &RepositorySQL[entity.Content]{db: db}
}

func (r *RepositorySQL[T]) Create(ctx context.Context, e entity.Content, args ...any) (string, error) {
	var lastID string
	q := postgres.BuildInsertQuery(tableName, []string{"name", "created_at", "updated_at"}, "RETURNING id")
	err := r.db.QueryRowContext(ctx, q, e.Name, e.CreatedAt, e.UpdatedAt).Scan(&lastID)
	if err != nil {
		return lastID, err
	}
	return lastID, nil
}

func (r *RepositorySQL[T]) ReadMany(ctx context.Context, limit, offset int, args ...any) (*sql.Rows, error) {
	q := postgres.BuildSelectQuery(tableName, []string{}, []string{"is_deleted", "name"}, "LIMIT $2 OFFSET $3", "OR")
	rows, err := r.db.QueryContext(ctx, q, args[0].(bool), limit, offset)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *RepositorySQL[T]) ReadOne(ctx context.Context, id string, args ...any) *sql.Row {
	q := postgres.BuildSelectQuery(tableName, []string{}, []string{"id"}, "LIMIT $2")
	return r.db.QueryRow(q, id, 1)
}

func (r *RepositorySQL[T]) Update(ctx context.Context, id string, e entity.Content, args ...any) (int64, error) {
	q := postgres.BuildUpdateQuery(tableName, []string{"name", "updated_at"}, []string{"id"}, "")
	res, err := r.db.Exec(q, e.Name, e.UpdatedAt, id)
	if err != nil {
		return -1, err
	}
	return postgres.GetRowsAffected(res), nil
}

func (r *RepositorySQL[T]) Delete(ctx context.Context, id string, args ...any) (int64, error) {
	q := postgres.BuildUpdateQuery(tableName, []string{"is_archived", "updated_at"}, []string{"id"}, "")
	res, err := r.db.Exec(q, true, args[0].(int64), id)
	if err != nil {
		return -1, err
	}
	return postgres.GetRowsAffected(res), nil
}

func (r *RepositorySQL[T]) DeleteHard(ctx context.Context, id string, args ...any) (int64, error) {
	q := postgres.BuildDeleteQuery(tableName, []string{"id"}, "")
	res, err := r.db.Exec(q, id)
	if err != nil {
		return -1, err
	}
	return postgres.GetRowsAffected(res), nil
}

func (r *RepositorySQL[T]) DB() *sql.DB {
	return r.db
}

// createMany Batch inserts contents
func (r *Repository[T]) createMany(ctx context.Context, entities []entity.Content) error {
	ctx1 := context.Background()
	ctx, cancelFn := context.WithTimeout(ctx1, 20*time.Second)
	defer cancelFn()
	dbConn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	err = dbConn.Raw(func(driverConn any) error {
		if conn, ok := driverConn.(*pgxstdlib.Conn); ok {
			var rows [][]any
			for _, e := range entities {
				rows = append(rows, []any{e.Name, e.CreatedAt, e.UpdatedAt})
			}
			copyCount, err := conn.Conn().CopyFrom(
				context.Background(),
				pgx.Identifier{tableName},
				[]string{"name", "created_at", "updated_at"},
				pgx.CopyFromRows(rows),
			)
			if err != nil {
				return err
			}
			l := len(entities)
			if int(copyCount) != l {
				return fmt.Errorf("bulk insert failed, insert count: %d param count: %d", copyCount, l)
			}
			return nil
		}
		return errors.New("driver connection is not of expected type")
	})
	if err != nil {
		return err
	}
	return nil
}
