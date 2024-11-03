package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type XKCD struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Image      string `json:"img"`
	Title      string `json:"title"`
	Day        string `json"day"`
}

func main() {
	var comics []XKCD
	for i := 0; i < 3; i++ {
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			log.Fatalf("Error getting file: %v", err)
		}
		fmt.Println(resp.Body)
		if resp.StatusCode == 200 {
			var comic XKCD
			json.NewDecoder(resp.Body).Decode(&comic)
			comics = append(comics, comic)
		}
	}
	jsonData, err := json.Marshal(comics)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("/Users/del/dev/gopl/ch4/xkcd/xkcd.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

/*
	dir := "/Users/del/dev/gopl/ch4/xkcd/comics"
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 120; i++ {
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		fileName := fmt.Sprintf("xkcd-%d", i)
		out, err := os.Create(dir + "/" + fileName + ".txt")
		if err != nil {
			log.Fatalf("Error creating file: %v", err)
		}
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Error getting url. Status: %s, Response: %s", resp.StatusCode, resp.Body)
		}
		defer resp.Body.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Fatalf("Error copying content: %v", err)
		}
*/
