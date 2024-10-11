// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", "_", "separator")
	flag.Parse()

	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
