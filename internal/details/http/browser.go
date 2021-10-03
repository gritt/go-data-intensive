package http

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type (
	HTTPClient interface {
		Get(url string) (*http.Response, error)
	}

	Browser struct {
		Client HTTPClient
	}
)

func NewBrowser(client HTTPClient) *Browser {
	return &Browser{
		Client: client,
	}
}

func (b Browser) GetContents(url url.URL) (string, error) {
	resp, err := b.Client.Get(url.String())
	if err != nil {
		return "", fmt.Errorf("request error: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	if resp.ContentLength == 0 {
		return "", errors.New("empty body: 0")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read error: %w", err)
	}

	return string(data), nil
}
