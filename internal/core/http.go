package core

import "net/url"

type HTTP interface {
	GetPageContents(url url.URL) (string, error)
}
