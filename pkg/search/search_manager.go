package search

import (
	"context"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"golang.org/x/net/html"
)

type SearchManager struct {
	ctx     context.Context
	cancel  context.CancelFunc
	s       Searcher
	resHtml string
}

func NewSearchManager(s Searcher, headless bool) SearchManager {
	var cancelFuncs []context.CancelFunc
	ctx, cancel := chromedp.NewExecAllocator(context.Background(),
		chromedp.Flag("headless", headless), chromedp.WindowSize(1920, 1080))
	cancelFuncs = append(cancelFuncs, cancel)
	ctx, cancel = context.WithTimeout(ctx, time.Second*30)
	cancelFuncs = append(cancelFuncs, cancel)
	ctx, cancel = chromedp.NewContext(ctx)
	cancelFuncs = append(cancelFuncs, cancel)

	cancel = func() {
		for _, c := range cancelFuncs {
			c()
		}
	}
	return SearchManager{
		ctx:    ctx,
		cancel: cancel,
		s:      s,
	}
}

func (sm *SearchManager) Search(query string) error {
	var resHTML string
	defer sm.cancel()
	err := chromedp.Run(
		sm.ctx,
		chromedp.Navigate(sm.s.Website()),
		chromedp.Sleep(time.Second),
		chromedp.Click(sm.s.SearchXPath()),
		chromedp.Sleep(time.Second),
		chromedp.SendKeys(sm.s.SearchXPath(), query+"\n"),
		chromedp.Sleep(time.Second*5),
		chromedp.InnerHTML(sm.s.ResultXPath(), &resHTML),
	)
	sm.resHtml = resHTML
	return err
}

func (sm *SearchManager) GetResults() (results []Result, err error) {
	root, err := html.Parse(strings.NewReader(sm.resHtml))
	if err != nil {
		return
	}
	return sm.parse(root), nil
}

func (sm *SearchManager) parse(n *html.Node) (results []Result) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		m := make(map[string]string)
		for _, a := range c.Attr {
			m[a.Key] = a.Val
		}
		if c.Data == "a" && sm.s.IsCorrectLink(c, m) {
			link, title := sm.s.ExtractLinkAndTitle(c, m)
			results = append(results, Result{Title: title, Link: link})
		}

		results = append(results, sm.parse(c)...)
	}
	return results
}
