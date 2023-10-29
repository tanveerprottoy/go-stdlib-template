package entity

import "database/sql"

type Content struct {
	ID        string `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	CreatedAt int64  `db:"created_at" json:"createdAt"`
	UpdatedAt int64  `db:"updated_at" json:"updatedAt"`
}

func (c *Content) ScanRows(rows *sql.Rows) error {
	return nil
}

func (c *Content) ScanRow(row *sql.Row) error {
	return nil
}