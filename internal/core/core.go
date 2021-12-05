package core

import (
	"context"
	"net/url"
)

// TODO, if it changes it breaks
// Statuses for messages
const (
	JOB_PENDING = iota
	JOB_DONE
	JOB_ERROR
)

/*
Core domain types and services
*/

type (
	// Message is an unprocessed webpage
	Message struct {
		Data   string
		UUID   string
		Status int
	}

	// WebPage is a core domain entity
	WebPage struct {
		UUID       int
		ParentUUID int
		Name       string
		URL        url.URL
		Links      []url.URL
	}

	// Searcher find webpages by given arguments
	Searcher interface {
		Search(ctx context.Context, args string) ([]WebPage, error)
	}

	// Indexer process raw data to webpages and stores it
	Indexer interface {
		Process(ctx context.Context, url string) (WebPage, error)
		Store(ctx context.Context, webpage WebPage) error
	}

	// Abstracts HTML parsing
	WebPageParser interface {
		Parse(ctx context.Context, data string) (WebPage, error)
	}

	// Abstracts integration with http
	Browser interface {
		GetContents(ctx context.Context, url url.URL) (string, error)
	}

	// Abstracts integration with a queue/stream/messaging system
	Messenger interface {
		Read(ctx context.Context, limit int) ([]Message, error)
		Write(ctx context.Context, msg Message) error
	}
)
