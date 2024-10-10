// Fetch prints the content found at a URL.
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		var buf bytes.Buffer
		b, err := io.Copy(&buf, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("number of bytes copied: ", b)
		var totalBuf int
		for i := range buf {
			totalBuf++
		}
		fmt.Println(totalBuf)
	}
}
