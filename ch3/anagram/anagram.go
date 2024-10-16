package main

import "fmt"

func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	strMap := make(map[string]int)
	for _, v := range s1 {
		strMap[string(v)]++
	}
	for _, v := range s2 {
		strMap[string(v)]--
	}
	for _, v := range strMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Anagram("wew", "wew"))
}
