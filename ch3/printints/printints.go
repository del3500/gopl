// ch3 exercises
package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
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

func FractionPart(s string) (string, error) {
	var buf bytes.Buffer
	i := strings.Index(s, ".")
	if i < 0 {
		return "", errors.New("no fraction part")
	}
	buf.WriteString(s[i:])
	return buf.String(), nil
}

func IntegerPart(s string) string {
	var buf bytes.Buffer
	i := strings.Index(s, ".")
	buf.WriteString(s[:i])
	return buf.String()
}

func AddComma(integer, fraction string) string {
	var buf bytes.Buffer
	n := len(integer)
	if n <= 3 {
		buf.WriteString(integer + fraction)
		return buf.String()
	}
	prefix := n % 3
	if prefix > 0 {
		buf.WriteString(integer[:prefix])
	}
	for i := prefix; i < n; i += 3 {
		if buf.Len() > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(integer[i : i+3])
	}
	buf.WriteString(fraction)
	return buf.String()
}

/* Bytes will take a string a parameter, the will add a comma
// every after 3rd character of the string starting from len(n - 1)
func BytesComma(s string) string {
	var buf bytes.Buffer
	var fBuf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	// Get the "." for to determine if it's floating point number.
	var dotOffset int
	for i := 0; i < n; i++ {
		if s[i] == '.' {
			dotOffset = i
			break
		}
	}
	if dotOffset > 0 {
		fBuf.WriteString(s[dotOffset:])
	}
	// calculate the number of characters before the first comma
	prefix := buf.Len() % 3
	if prefix > 0 {
		buf.WriteString(s[:prefix])
	}
	for i := prefix; i < n; i += 3 {
		if buf.Len() > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i : i+3])
	}
	buf.Write(fBuf.Bytes())
	return buf.String()
}*/

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}
