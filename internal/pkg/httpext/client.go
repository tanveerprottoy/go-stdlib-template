package httpext

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

type ClientProvider struct {
	HTTPClient *http.Client
}

func NewClientProvider(timeout time.Duration, transport *http.Transport, checkRedirectFunc func(req *http.Request, via []*http.Request) error) *ClientProvider {
	c := new(ClientProvider)
	c.HTTPClient = &http.Client{
		Timeout: timeout,
	}
	if transport != nil {
		c.HTTPClient.Transport = transport
	}
	if checkRedirectFunc != nil {
		c.HTTPClient.CheckRedirect = checkRedirectFunc
	}
	return c
}

func (c *ClientProvider) Request(method string, url string, header http.Header, body io.Reader) (int, io.ReadCloser, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return -1, nil, err
	}
	if header != nil {
		req.Header = header
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return -1, nil, err
	}
	return res.StatusCode, res.Body, nil
}

// ex:
// resp, err := http.PostForm("http://example.com/form",
// url.Values{"key": {"Value"}, "id": {"123"}})
func (c *ClientProvider) PostForm(url string, header http.Header, values url.Values) (int, io.ReadCloser, error) {
	res, err := http.PostForm(url, values)
	if err != nil {
		return -1, nil, err
	}
	return res.StatusCode, res.Body, nil
}
