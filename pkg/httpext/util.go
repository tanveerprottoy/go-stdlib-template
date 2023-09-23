package httpext

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/jsonext"
)

func GetURLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func ParseAuthToken(r *http.Request) ([]string, error) {
	tkHeader := r.Header.Get("Authorization")
	if tkHeader == "" {
		// Token is missing
		return nil, errors.New("auth token is missing")
	}
	splits := strings.Split(tkHeader, " ")
	// token format is `Bearer {tokenBody}`
	if len(splits) != 2 {
		return nil, errors.New("token format is invalid")
	}
	return splits, nil
}

// ParseRequestBody parses the request body
// The caller must pass the address for the v any param, ex: &v
func ParseRequestBody(r io.ReadCloser, v any) error {
	defer r.Close()
	err := jsonext.Decode(r, v)
	if err != nil {
		return err
	}
	return nil
}

func BuildURL(base, path string, queriesMap map[string]string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	if path != "" {
		// Path param
		u.Path += path
	}
	if queriesMap != nil {
		// Query params
		p := url.Values{}
		for k, v := range queriesMap {
			p.Add(k, v)
		}
		u.RawQuery = p.Encode()
	}
	return u.String(), nil
}
