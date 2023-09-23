package postgres

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/errorext"
)

// ScanRows convert rows to struct slice
// must provide pointer address for params
func ScanRows[T any](rows *sql.Rows, e *T, params ...any) ([]T, errorext.HTTPError) {
	d := []T{}
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(params...); err != nil {
			return nil, errorext.BuildDBError(err)
		}
		d = append(d, *e)
	}
	return d, errorext.HTTPError{}
}

// ScanRow convert row to struct
// must provide pointer address for params
func ScanRow[T any](row *sql.Row, obj *T, params ...any) errorext.HTTPError {
	if err := row.Scan(params...); err != nil {
		return errorext.BuildDBError(err)
	}
	return errorext.HTTPError{}
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

// ScanRowsBasic scans rows
// must pass pointer address for e param
func ScanRowsBasic[T any](rows *sql.Rows, e *T) errorext.HTTPError {
	if rows.Next() {
		err := rows.Scan(e)
		if err != nil {
			return errorext.HTTPError{Code: http.StatusInternalServerError, Err: err}
		}
	}
	return errorext.HTTPError{}
}

// ScanRowBasic scans rows
// must pass pointer address for e param
func ScanRowBasic[T any](row *sql.Row, e *T) errorext.HTTPError {
	err := row.Scan(e)
	if err != nil {
		return errorext.HTTPError{Code: http.StatusInternalServerError, Err: err}
	}
	return errorext.HTTPError{}
}

func GetRowsAffected(result sql.Result) int64 {
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	return rows
}
