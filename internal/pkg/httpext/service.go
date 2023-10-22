package httpext

import (
	"io"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/jsonext"
)

func Request[T any](method string, url string, header http.Header, body io.Reader, client *ClientProvider) (*T, map[string]any, error) {
	code, resBody, err := client.Request(method, url, header, body)
	if err != nil {
		return nil, nil, err
	}
	if code >= http.StatusOK && code < http.StatusMultipleChoices {
		// res ok, parse response body to type
		var d T
		err := jsonext.Decode(resBody, &d)
		if err != nil {
			return nil, nil, err
		}
		return &d, nil, nil
	} else {
		// res not ok, parse error
		var errRes map[string]any
		err := jsonext.Decode(resBody, &errRes)
		if err != nil {
			return nil, nil, err
		}
		return nil, errRes, nil
	}
}
