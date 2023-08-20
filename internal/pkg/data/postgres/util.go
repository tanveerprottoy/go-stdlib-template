package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/adapter"
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
func GetEntities[T any](rows *sql.Rows, obj *T, params ...any) ([]T, error) {
	var e []T
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(params...); err != nil {
			return nil, fmt.Errorf("GetEntities %v", err)
		}
		adapter.ValuesToStruct(params, obj)
		e = append(e, *obj)
	}
	return e, nil
}

func GetEntity[T any](row *sql.Row, obj *T, params ...any) (T, error) {
	if err := row.Scan(params...); err != nil {
		return *obj, fmt.Errorf("GetEntity %v", err)
	}
	adapter.ValuesToStruct(params, obj)
	return *obj, nil
}
