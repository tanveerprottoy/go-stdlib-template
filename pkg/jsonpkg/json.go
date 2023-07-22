package jsonpkg

import (
	"encoding/json"
	"io"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(d []byte, v any) error {
	return json.Unmarshal(d, &v)
}

func Encode(v any, w io.Writer) error {
	return json.NewEncoder(w).Encode(v)
}

func Decode(v any, r io.Reader) error {
	return json.NewDecoder(r).Decode(&v)
}
