// ch3 exercises
package main

import (
	"bytes"
	"fmt"
)

// intsToString is like fmt.Sprintf(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteRune('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteRune(']')
	return buf.String()
}

// Bytes will take a string a parameter, the will add a comma
// every after 3rd character of the string starting from len(n - 1)
func BytesComma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	// calculate the number of characters before the first comma
	prefix := n % 3
	if prefix > 0 {
		buf.WriteString(s[:prefix])
	}
	for i := prefix; i < n; i += 3 {
		if buf.Len() > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
	fmt.Println(BytesComma("1231232134"))
}
