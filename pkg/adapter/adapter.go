package adapter

import (
	"errors"
	"io"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/jsonext"
)

func IOReaderToBytes(r io.Reader) ([]byte, error) {
	b, err := io.ReadAll(r)
	return b, err
}

func BytesToType[T any](b []byte) (*T, error) {
	var out T
	err := jsonext.Unmarshal(b, &out)
	return &out, err
}

func BodyToType[T any](r io.ReadCloser) (*T, error) {
	var out T
	err := jsonext.Decode(r, &out)
	if err != nil {
		return nil, err
	}
	return AnyToType[T](out)
}

func AnyToType[T any](v any) (*T, error) {
	var out T
	b, err := jsonext.Marshal(v)
	if err != nil {
		return nil, err
	}
	err = jsonext.Unmarshal(b, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func InterfaceToStruct[T any](inter interface{}) (T, error) {
	s, ok := inter.(T)
	if ok {
		return s, errors.New("TypeCast error")
	}
	return s, nil
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringToFloat(s string, bitSize int) (float64, error) {
	return strconv.ParseFloat(s, bitSize)
}

// ValuesToStruct will set the values provided on the
// struct provided in the param
// caller must provide pointer addresses of values and t
func ValuesToStruct[T any](values []any, obj T) {
	v := reflect.Indirect(
		reflect.ValueOf(obj).Elem(),
	)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.CanSet() {
			param := values[i]
			switch f.Kind() {
			case reflect.String:
				f.SetString(
					reflect.ValueOf(param).Elem().Interface().(string),
				)
			case reflect.Int32, reflect.Int64:
				f.SetInt(reflect.ValueOf(param).Elem().Interface().(int64))
			case reflect.Float32, reflect.Float64:
				f.SetFloat(reflect.ValueOf(param).Elem().Interface().(float64))
			case reflect.Bool:
				f.SetBool(reflect.ValueOf(param).Elem().Interface().(bool))
			case reflect.Struct:
				// currently only handle time.Time type
				f.Set(reflect.ValueOf(
					reflect.ValueOf(param).Elem().Interface().(time.Time),
				))
			default:
				log.Println("type unknown")
			}
		}
	}
}
