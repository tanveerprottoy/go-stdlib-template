package errorpkg

import (
	"errors"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/constant"
)

func NewError(m string) error {
	return errors.New(m)
}

func MakeHTTPError(code int, err error) HTTPError {
	return HTTPError{code, err}
}

func HandleDBError(err error) *HTTPError {
	httpErr := &HTTPError{Code: http.StatusBadRequest, Err: err}
	if err.Error() == "sql: no rows in result set" {
		httpErr.Code = http.StatusNotFound
		httpErr.Err = NewError(constant.NotFound)
	}
	return httpErr
}
