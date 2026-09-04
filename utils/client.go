package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

var defaultClient = &http.Client{
	Timeout: 30 * time.Second,
}

func Request(method, url string, data ...byte) (*http.Response, error) {
	var datareader io.Reader
	if len(data) > 0 {
		datareader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, url, datareader)
	if err != nil {
		return nil, err
	}

	//	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %v %v", resp.StatusCode, resp.Status)
	}

	return resp, nil

}

func Get(url string) ([]byte, error) {
	resp, err := Request(http.MethodGet, url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
