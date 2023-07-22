package search

import "golang.org/x/net/html"

type YoutubeSearcher struct{}

func (g YoutubeSearcher) Website() string {
	return "https://www.youtube.com"
}
func (g YoutubeSearcher) SearchXPath() string {
	return "/html/body/ytd-app/div[1]/div/ytd-masthead/div[4]/div[2]/ytd-searchbox/form/div[1]/div[1]/input"
}
func (g YoutubeSearcher) ResultXPath() string {
	return `/html/body/ytd-app/div[1]/ytd-page-manager/ytd-search/div[1]/ytd-two-column-search-results-renderer/div/ytd-section-list-renderer/div[2]/ytd-item-section-renderer`
}

func (g YoutubeSearcher) IsCorrectLink(linkNode *html.Node, attrs map[string]string) bool {
	return attrs["id"] == "video-title"
}

func (g YoutubeSearcher) ExtractLinkAndTitle(linkNode *html.Node, attrs map[string]string) (link, title string) {
	return g.Website() + attrs["href"], attrs["title"]
}
