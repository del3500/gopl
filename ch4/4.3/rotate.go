package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5}
	reverse(s, 0, len(s)-1)
	fmt.Println(s)

}

func reverse(s []int, start, end int) {
	for ; start < end; start, end = start+1, end-1 {
		s[start], s[end] = s[end], s[start]
	}
}

// TODO: rotate right by k times.
func rotateRight(s []int, k int) {

}
