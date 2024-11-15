package main

import (
	links "delmesia/gopl/ch5/anonymous_functions"
	"fmt"
	"log"
	"os"
)

func main() {
	links.BreadthFirst(crawl, os.Args[1:])
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
