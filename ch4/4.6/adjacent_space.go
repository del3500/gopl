package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []byte("hello,      world")
	squash(&s)
	fmt.Println(string(s))

}

func squash(data *[]byte) {
	if len(*data) == 0 {
		return
	}
	// pointer for the position to write the next character
	j := 0
	spaceFound := false

	for i := 0; i < len(*data); i++ {
		if unicode.IsSpace(rune((*data)[i])) {
			// only write a space if we haven't written one yet
			if !spaceFound {
				// write a single ASCII space
				(*data)[j] = ' '
				j++
				spaceFound = true
			}
		} else {
			// write the current character
			(*data)[j] = (*data)[i]
			j++
			// reset the found flag
			spaceFound = false
		}
	}
	// resize the slice to new length
	(*data) = (*data)[:j]
}
