package validator

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/errorpkg"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/jsonpkg"
)

// ParseValidateRequestBody parses and validates the request body
// The caller must pass the address for the v any param, ex: &v
func ParseValidateRequestBody(r io.ReadCloser, v any, validate *validator.Validate) ([]errorpkg.ValidationError, error) {
	defer r.Close()
	var validationErrs []errorpkg.ValidationError
	err := jsonpkg.Decode(r, v)
	if err != nil {
		return nil, err
	}
	// validate request body
	err = validate.Struct(v)
	if err != nil {
		// Struct is invalid
		var v errorpkg.ValidationError
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
			v.Message = err.Field() + " " + err.Tag()
			validationErrs = append(validationErrs, v)
		}
	}
	return validationErrs, err
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
