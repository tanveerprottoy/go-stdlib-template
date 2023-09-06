package httpext

type ErrorBody struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type RemoteErrorBody struct {
	Errors string `json:"errors"`
}
