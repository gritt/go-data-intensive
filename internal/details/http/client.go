package http

import "net/url"

type Client struct {
}

func (c Client) GetPageContents(url url.URL) (string, error) {
	return "", nil
}
