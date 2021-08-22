package core

type WebPageParser interface {
	Parse(data string) (WebPage, error)
}
