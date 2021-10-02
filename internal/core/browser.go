package core

import "net/url"

type Browser interface {
	GetContents(url url.URL) (string, error)
}
