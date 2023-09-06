package errorext

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/lib/pq"
)

func NewError(msg string) error {
	return errors.New(msg)
}

func BuildDBError(err any) HTTPError {
	fmt.Println("BuildDBError.err: ", err)
	httpErr := HTTPError{Code: http.StatusInternalServerError, Err: NewError("internal server error")}
	if pqErr, ok := err.(*pq.Error); ok {
		fmt.Println("BuildDBError.err: ", pqErr)
		switch pqErr.Code {
		case SQLCodeNotFound:
			// "20000": "case_not_found",
			httpErr.Code = http.StatusNotFound
			httpErr.Err = NewError("not found")
		case SQLCodeNoData:
			// "02000": "no data",
			httpErr.Code = http.StatusNotFound
			httpErr.Err = NewError("not found")
		case SQLCodeInvalidTextRepresentation:
			// "22P02": "invalid_text_representation",
			httpErr.Code = http.StatusBadRequest
			httpErr.Err = NewError("the data provided is invalid")	
		}
	}
	// check if it's an error type
	if err, ok := err.(error); ok {
		fmt.Print(err)
		switch err {
		case sql.ErrNoRows:
			httpErr.Code = http.StatusNotFound
			httpErr.Err = NewError("not found")
		}
	}
	return httpErr
}
