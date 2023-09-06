package errorext

type HTTPError struct {
	Code int
	Err  error
}

type ValidationError struct {
	Message string `json:"message"`
}

const (
	// case_not_found
	SQLCodeNotFound = "20000"
	// no_data
	SQLCodeNoData = "02000"
	// invalid_text_representation
	SQLCodeInvalidTextRepresentation = "22P02"
)
