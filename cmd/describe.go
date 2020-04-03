package cmd

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

func describe(id string) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")

	for _, element := range feed.Items {
		if element.GUID == id {
			fmt.Println(element.Title)
			fmt.Println(" ")
			fmt.Println(element.Description)
			return
		}
	}
	fmt.Println("Article " + id + " not found.")
}
