package main

import "fmt"

func main() {
	fmt.Println(sum(1, 1, 1, 1))
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
