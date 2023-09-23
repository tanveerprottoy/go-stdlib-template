package validatorext

import (
	"io"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/jsonext"
)

// ValidateStruct parses and validates the request body
// The caller must pass the address for the param v, ex: &v
func ValidateStruct(v any, validate *validator.Validate) []string {
	var validationErrs []string
	// validate request body
	err := validate.Struct(v)
	if err != nil {
		// Struct is invalid
		var msg string
		for _, err := range err.(validator.ValidationErrors) {
			msg = err.Field() + " " + err.Tag()
			validationErrs = append(validationErrs, msg)
		}
	}
	return validationErrs
}

// ParseValidateRequestBody parses and validates the request body
// The caller must pass the address for the v any param, ex: &v
func ParseValidateRequestBody(r io.ReadCloser, v any, validate *validator.Validate) ([]string, error) {
	defer r.Close()
	var validationErrs []string
	err := jsonext.Decode(r, v)
	if err != nil {
		return nil, err
	}
	// validate request body
	err = validate.Struct(v)
	if err != nil {
		// Struct is invalid
		var msg string
		for _, err := range err.(validator.ValidationErrors) {
			msg = err.Field() + " " + err.Tag()
			validationErrs = append(validationErrs, msg)
		}
	}
	return validationErrs, err
}

// RegisterTagNameFunc configures validator to use
// defined json name to use as struct field name
func RegisterTagNameFunc(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		n := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if n == "-" {
			return ""
		}
		return n
	})
}

func NotEmpty(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return len(strings.TrimSpace(field.String())) > 0
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}
