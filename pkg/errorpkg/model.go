package errorpkg

type HTTPError struct {
	Code int
	Err  error
}

type ValidationError struct {
	Message string `json:"message"`
}
