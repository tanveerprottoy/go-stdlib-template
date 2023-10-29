package postgres

import "database/sql"

type Scanner interface {
	ScanRows(rows *sql.Rows) error

	ScanRow(row *sql.Row) error
}

type Row interface {
	Scan(...interface{}) error
}
