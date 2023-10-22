package errorext

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
)
const (
	// case_not_found
	SQLCodeNotFound = "20000"
	// no_data
	SQLCodeNoData = "02000"
	// invalid_text_representation
	SQLCodeInvalidTextRepresentation = "22P02"
	// undefined_function
	SQLCodeUndefinedFunction = "42883"
	// undefined_table
	SQLCodeUndefinedTable = "42P01"
	// undefined_parameter
	SQLCodeUndefinedParam = "42P02"
	// invalid_column_reference
	SQLInvalidColumnReference = "42P10"
)

func BuildDBError(err error) HTTPError {
	httpErr := HTTPError{Code: http.StatusInternalServerError, Err: errors.New("internal server error")}
	// check if it's an sql error
	switch err {
	case sql.ErrNoRows:
		httpErr.Code = http.StatusNotFound
		httpErr.Err = errors.New("not found")
		return httpErr
	case sql.ErrTxDone:
		httpErr.Code = http.StatusNotFound
		httpErr.Err = errors.New("transaction already closed")
		return httpErr
	}
	var pgErr *pgconn.PgError
	if ok := errors.As(err, &pgErr); ok {
		switch pgErr.Code {
		case SQLCodeNotFound:
			// "20000": "case_not_found",
			httpErr.Code = http.StatusNotFound
			httpErr.Err = errors.New("not found")
			return httpErr
		case SQLCodeNoData:
			// "02000": "no data",
			httpErr.Code = http.StatusNotFound
			httpErr.Err = errors.New("not found")
			return httpErr
		case SQLCodeInvalidTextRepresentation:
			// "22P02": "invalid_text_representation",
			// potentially need to log this
			httpErr.MainErr = pgErr
			httpErr.Code = http.StatusBadRequest
			httpErr.Err = errors.New("the data provided is invalid")
			return httpErr
		case SQLCodeUndefinedFunction:
			// "42883": "undefined_function",
			// potentially need to log this
			httpErr.MainErr = pgErr
			httpErr.Code = http.StatusInternalServerError
			httpErr.Err = errors.New("the expected resource is not available")
			return httpErr
		case SQLCodeUndefinedTable:
			// potentially need to log this
			httpErr.MainErr = pgErr
			httpErr.Code = http.StatusInternalServerError
			httpErr.Err = errors.New("the expected resource is not available")
			return httpErr
		case SQLCodeUndefinedParam:
			// potentially need to log this
			httpErr.MainErr = pgErr
			httpErr.Code = http.StatusInternalServerError
			httpErr.Err = errors.New("the expected resource is not available")
			return httpErr
		case SQLInvalidColumnReference:
			// potentially need to log this
			httpErr.MainErr = pgErr
			httpErr.Code = http.StatusInternalServerError
			httpErr.Err = errors.New("the expected resource is not available")
			return httpErr
		}
	}
	return httpErr
}
