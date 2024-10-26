// Issues prints a table of Github issues matching the search terms
package main

import (
	"delmesia/gopl/ch4/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		tn := time.Now()
		formattedTime := item.CreatedAt.Format("2006-01-02 15:04:05")
		layout := "2006-01-02 15:04:05"
		parsedTime, err := time.Parse(layout, formattedTime)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return
		}
		months := int(tn.Year()-parsedTime.Year())*12 + int(tn.Month()-parsedTime.Month())
		fmt.Printf("%-3d months old #%-5d %9.9s %.55s\n", months, item.Number, item.User.Login, item.Title)
	}
}
