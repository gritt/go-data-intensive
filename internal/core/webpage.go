package core

import "net/url"

type WebPage struct {
	ID       int
	Name     string
	URL      url.URL
	Links    []url.URL
	Keywords []string
}
