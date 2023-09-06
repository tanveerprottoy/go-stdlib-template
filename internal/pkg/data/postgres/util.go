package postgres

import (
	"database/sql"
	"log"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/errorext"
)

func GetRowsAffected(result sql.Result) int64 {
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	return rows
}

// GetEntities convert rows to entity slice
// must provide pointer address for params
func GetEntities[T any](rows *sql.Rows, obj *T, params ...any) ([]T, errorext.HTTPError) {
	d := []T{}
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(params...); err != nil {
			return nil, errorext.BuildDBError(err)
		}
		d = append(d, *obj)
	}
	return d, errorext.HTTPError{}
}

func GetEntity[T any](row *sql.Row, obj *T, params ...any) (T, errorext.HTTPError) {
	if err := row.Scan(params...); err != nil {
		return *obj, errorext.BuildDBError(err)
	}
	return *obj, errorext.HTTPError{}
}
