package httpUtils

import (
	"io"
	"net/http"
	"time"
)

func ForwardRequest(method, url string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return resp, err
	}
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
