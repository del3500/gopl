package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Movie struct {
	Title    string
	Director string
	Actors   []string
	Showing  bool
}

func main() {
	m := []Movie{
		{
			Title:    "sample",
			Director: "idk",
			Actors:   []string{"ck", "batman", "ming"},
			Showing:  true,
		},
		{
			Title:    "sample2",
			Director: "idk2",
			Actors:   []string{"pia"},
			Showing:  false,
		},
	}
	f, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewEncoder(f).Encode(m); err != nil {
		log.Fatal(err)
	}

	var titles []struct{ Title string }
	jf, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewDecoder(jf).Decode(&titles); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", titles)
}
