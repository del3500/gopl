package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := `<html><body><p>Hello world!</p><img src="image.jpg" /><p>Another a a paragraph.</p></body></html>`
	doc, err := html.Parse(strings.NewReader(url))
	if err != nil {
		panic(err)
	}
	words, images := countWordsAndImages(doc)
	fmt.Printf("words: %d, images: %d", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HMTL: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	if n.Type == html.ElementNode && n.Data == "p" {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			words += countWords(c.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childWords, childImages := countWordsAndImages(c)
		words += childWords
		images += childImages
	}
	return
}

func countWords(txt string) int {
	words := strings.Fields(txt)
	return len(words)
}
