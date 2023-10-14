package postgres

import "strconv"

func BuildInsertQuery(tableName string, columns []string, clause string) string {
	q := "INSERT INTO " + tableName
	cols := " ("
	vals := " VALUES ("
	for i, v := range columns {
		cols += v + ", "
		vals += "$" + strconv.Itoa(i+1) + ", "
	}
	// remove trailing space & comma
	// there are two chars to be removed
	cols = cols[:len(cols)-2]
	vals = vals[:len(vals)-2]
	// add closing parentheses
	cols += ")"
	vals += ")"
	return q + cols + vals + " " + clause
}

func BuildSelectQuery(tableName string, projections []string, whereClauseCols []string, clause string) string {
	q := "SELECT "
	p := ""
	w := ""
	if len(projections) > 0 {
		for _, v := range projections {
			p += v + ", "
		}
		// remove trailing space & comma
		// there are two chars to be removed
		p = p[:len(p)-2]
	} else {
		p += "*"
	}
	// where clause length
	lenWhereClause := len(whereClauseCols)
	if lenWhereClause > 0 {
		w += " WHERE "
		for i, v := range whereClauseCols {
			w += v + " = $" + strconv.Itoa(i+1) + ", "
		}
		// remove trailing space & comma
		// there are two chars to be removed
		w = w[:len(w)-2]
	}
	return q + p + " FROM " + tableName + w + " " + clause
}

func BuildUpdateQuery(tableName string, columns []string, whereClauseCols []string, clause string) string {
	q := "UPDATE " + tableName
	c := " SET "
	w := " WHERE "
	// col length
	lenCols := len(columns)
	if lenCols > 0 {
		for i, v := range columns {
			c += v + " = $" + strconv.Itoa(i+1) + ", "
		}
		// remove trailing space & comma
		// there are two chars to be removed
		c = c[:len(c)-2]
	}
	for i, v := range whereClauseCols {
		w += v + " = $" + strconv.Itoa(i+1+lenCols) + ", "
	}
	// remove trailing space & comma
	// there are two chars to be removed
	w = w[:len(w)-2]
	return q + c + w + " " + clause
}

func BuildDeleteQuery(tableName string, whereClauseCols []string, clause string) string {
	q := "DELETE FROM " + tableName
	w := " WHERE "
	for i, v := range whereClauseCols {
		w += v + " = $" + strconv.Itoa(i+1) + ", "
	}
	// remove trailing space & comma
	// there are two chars to be removed
	w = w[:len(w)-2]
	return q + w + " " + clause
}
