package core

import (
	"context"
)

type SearchEngine struct {
}

func NewSearchEngine() *SearchEngine {
	return &SearchEngine{}
}

func (s *SearchEngine) Search(ctx context.Context, args string) ([]WebPage, error) {
	// fmt.Println("SearchEngine Search...")
	return []WebPage{}, nil
}
