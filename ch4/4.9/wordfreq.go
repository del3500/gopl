package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freq := make(map[string]int)
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		letter := scanner.Text()
		freq[letter]++
	}

	fmt.Printf("words\tcount\n")
	for k, v := range freq {
		fmt.Printf("%s\t%d\n", k, v)
	}

}
