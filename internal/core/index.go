package core

import (
	"context"
)

type SearchIndex struct {
	browser Browser
	parser  WebPageParser
}

func NewSearchIndex() *SearchIndex {
	return &SearchIndex{}
}

func (s *SearchIndex) Process(ctx context.Context, url string) (WebPage, error) {
	// fmt.Println("SearchIndex Process...")
	return WebPage{}, nil
}

func (s *SearchIndex) Store(ctx context.Context, webpage WebPage) error {
	// fmt.Println("SearchIndex Store...")
	return nil
}
