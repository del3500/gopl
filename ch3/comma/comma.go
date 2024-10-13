// Implementation of Comma
package main

// Comma inserts commas in non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3] + "," + s[n-3:])
}
