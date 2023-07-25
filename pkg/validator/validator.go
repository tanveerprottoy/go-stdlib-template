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

func ParseValidateRequestBody[T any](reader io.Reader, typ T, validate *validator.Validate) (T, error, []errorpkg.ValidationError) {
	var vErrs []errorpkg.ValidationError
	err := jsonpkg.Decode(typ, reader)
	if err != nil {
		return typ, err, nil
	}
	// validate request body
	err = validate.Struct(typ)
	if err != nil {
		// Struct is invalid
		var v errorpkg.ValidationError
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
			v.Message = err.Field() + " " + err.Tag()
			vErrs = append(vErrs, v)
		}
	}
	return typ, err, vErrs
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
