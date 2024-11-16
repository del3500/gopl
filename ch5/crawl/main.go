package main

import (
	links "delmesia/gopl/ch5/anonymous_functions"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	links.BreadthFirst(crawl, os.Args[1:])
}

func crawl(url string) []string {
	fmt.Println(url)
	lists, err := links.Extract(url)
	for i, list := range lists {
		resp, err := http.Get("https://" + list)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		file := fmt.Sprintf("%d", i)
		f, err := os.Create(file)
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err != nil {
		log.Print(err)
	}
	return lists
}
