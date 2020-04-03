package cmd

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

func list() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")

	for i := 0; i < 5; i++ {
		item := feed.Items[i]
		fmt.Println(item.Title + ": " + item.GUID)
	}
}
