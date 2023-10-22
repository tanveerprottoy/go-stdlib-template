package errorext

type HTTPError struct {
	Code    int
	Err     error
	MainErr error
}

func MakeHTTPError(code int, err error, mainErr error) HTTPError {
	return HTTPError{Code: code, Err: err, MainErr: mainErr}
}
