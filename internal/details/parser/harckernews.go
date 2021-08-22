package parser

import (
	"net/url"

	"github.com/gritt/go-data-intensive/internal/core"
)

type HackerNewsParser struct {
}

func (h HackerNewsParser) Parse(data string) (core.WebPage, error) {
	panic("implement me")

	return core.WebPage{
		ID:       0,
		ParentID: 0,
		Name:     "",
		URL:      url.URL{},
		Links:    nil,
		Keywords: nil,
	}, nil
}
