package main

import "fmt"

func main() {
	s := []string{"x", "x", "z", "y", "x", "x"}
	inplace(&s)
	fmt.Println(s)
}

func inplace(str *[]string) {
	if len(*str) == 0 {
		return
	}
	j := 1
	for i := 1; i < len(*str); i++ {
		if (*str)[i] != (*str)[i-1] {
			(*str)[j] = (*str)[i]
			j++
		}
	}
	*str = (*str)[:j]
}
