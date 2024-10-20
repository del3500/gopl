package main

import "fmt"

// Circular array

func main() {

	s := []int{1, 2, 3, 4, 5}
	rotateRight(s, 3)
	fmt.Println(s)

}

func reverse(s []int, start, end int) {
	for ; start < end; start, end = start+1, end-1 {
		s[start], s[end] = s[end], s[start]
	}
}

// TODO: rotate right by k times.
func rotateRight(s []int, k int) {
	n := len(s)
	if n == 0 {
		return
	}

	k = k % n
	if k == 0 {
		return
	}

	// Reverse the entire array
	reverse(s, 0, n-1)

	// Reverse the first k-elements
	reverse(s, 0, k-1)

	// Reverse the remaining n-k elements
	reverse(s, k, n-1)
}
