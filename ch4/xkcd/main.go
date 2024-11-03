package main

import (
	"encoding/json"
	"fmt"
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
	Day        string `json:"day"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("xkcd.json")
	check(err)
	defer f.Close()
	dec := json.NewDecoder(f)

	_, err = dec.Token()
	check(err)

	for dec.More() {
		var x XKCD
		err := dec.Decode(&x)
		check(err)
		fmt.Printf("%v: %v\n", x.SafeTitle, x.Year)
	}
	_, err = dec.Token()
	check(err)
}

/*func getComics(url string, ch chan<- XKCD, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		var comic XKCD
		json.NewDecoder(resp.Body).Decode(&comic)
		ch <- comic
	}
}*/

/*
	var comics []XKCD
	ch := make(chan XKCD)
	var wg sync.WaitGroup
	for i := 1; i < 3000; i++ {
		wg.Add(1)
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		go getComics(url, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	for comic := range ch {
		comics = append(comics, comic)
	}
	data, err := json.Marshal(comics)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("/Users/del/dev/gopl/ch4/xkcd/xkcd.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}*/
