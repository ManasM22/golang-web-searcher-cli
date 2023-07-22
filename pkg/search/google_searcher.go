package search

import "golang.org/x/net/html"

type GoogleSearcher struct{}

func (g GoogleSearcher) Website() string {
	return "https://www.google.com/"
}
func (g GoogleSearcher) SearchXPath() string {
	return "/html/body/div[1]/div[3]/form/div[1]/div[1]/div[1]/div/div[2]/textarea"
}
func (g GoogleSearcher) ResultXPath() string {
	return `//*[@id="rso"]`
}

func (g GoogleSearcher) IsCorrectLink(linkNode *html.Node, attrs map[string]string) bool {
	for _, a := range []string{"jscontroller", "jsname", "jsaction", "data-ved"} {
		if _, ok := attrs[a]; !ok {
			return false
		}
	}
	return true
}
func (g GoogleSearcher) ExtractLinkAndTitle(linkNode *html.Node, attrs map[string]string) (link, title string) {
	for c := linkNode.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "h3" {
			return attrs["href"], c.FirstChild.Data
		}
	}
	return attrs["href"], title
}
