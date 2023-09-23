package global

import (
	"context"
	"database/sql"

	"github.com/tanveerprottoy/starter-go/stdlib/pkg/data/sql/postgres"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/errorext"
)

type Count struct {
	Count int `db:"count" json:"count"`
}

func CalculateOffset(limit, page int) int {
	return limit * (page - 1)
}

func FetchAndScanTotalCount(tableName, colName, clause string, db *sql.DB, ctx context.Context) (Count, errorext.HTTPError) {
	row := GetTotalCount(tableName, colName, clause, db, ctx)
	var entity Count
	httpErr := postgres.ScanRow(row, &entity, &entity.Count)
	return entity, httpErr
}

func GetTotalCount(tableName, colName, clause string, db *sql.DB, ctx context.Context) *sql.Row {
	q := "SELECT * FROM get_total_row_count($1, $2, $3)"
	return db.QueryRow(q, tableName, colName, clause)
}