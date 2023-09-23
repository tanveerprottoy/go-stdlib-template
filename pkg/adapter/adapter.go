package adapter

import (
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/jsonext"
)

func IOReaderToBytes(r io.Reader) ([]byte, error) {
	b, err := io.ReadAll(r)
	return b, err
}

func BytesToType[T any](b []byte) (*T, error) {
	var out T
	err := json.Unmarshal(b, &out)
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
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &out)
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
