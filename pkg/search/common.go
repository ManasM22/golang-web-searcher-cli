package search

import "golang.org/x/net/html"

type Searcher interface {
	Website() string
	SearchXPath() string
	ResultXPath() string

	IsCorrectLink(linkNode *html.Node, attrs map[string]string) bool
	ExtractLinkAndTitle(linkNode *html.Node, attrs map[string]string) (link, title string)
}

type Result struct {
	Title, Link string
}
