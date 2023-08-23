package errorpkg

import (
	"errors"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/constant"
)

func NewError(msg string) error {
	return errors.New(msg)
}

func BuildDBError(err error) *HTTPError {
	httpErr := &HTTPError{Code: http.StatusBadRequest, Err: err}
	if err.Error() == "sql: no rows in result set" {
		httpErr.Code = http.StatusNotFound
		httpErr.Err = NewError(constant.NotFound)
	}
	return httpErr
}
