package main

import (
	"flag"
	"fmt"

	"youtube.com/pkg/search"
)

func main() {
	g := flag.Bool("google", false, "Search in google?")
	y := flag.Bool("youtube", false, "Search in youtube?")
	query := flag.String("query", "hello world", "what to search?")
	flag.Parse()

	ss := []search.Searcher{}
	if *g {
		ss = append(ss, search.GoogleSearcher{})
	}
	if *y {
		ss = append(ss, search.YoutubeSearcher{})
	}

	for _, s := range ss {
		sm := search.NewSearchManager(s, true)
		fmt.Println("    Search Results from", s.Website())
		if err := sm.Search(*query); err != nil {
			fmt.Println("Error:", err)
		}
		results, err := sm.GetResults()
		if err != nil {
			fmt.Println("Error:", err)
		}
		for i, r := range results {
			fmt.Printf("%d. %s\n", i+1, r.Title)
			fmt.Printf("%s\n", r.Link)
		}
	}
}
