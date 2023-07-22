package main

import (
	"fmt"
	"log"

	"youtube.com/pkg/search"
)

func main() {
	s := search.YoutubeSearcher{}
	sm := search.NewSearchManager(s, true)

	check(sm.Search("What is 2+2"))
	results, err := sm.GetResults()
	check(err)
	for i, r := range results {
		fmt.Printf("\n%d. %s\n", i+1, r.Title)
		fmt.Printf("%s\n", r.Link)
		fmt.Println("------------------------")
	}
}

func check(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
	}
}
