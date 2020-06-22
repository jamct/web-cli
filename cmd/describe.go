package cmd

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

func describe(id string) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")

	for _, element := range feed.Items {
		if element.GUID == id {
			fmt.Println(formatDate(element.Published))
			fmt.Println(element.Title)
			fmt.Println(" ")
			fmt.Println(element.Description)
			return
		}
	}
	fmt.Println("Article " + id + " not found.")
}

func formatDate(date string) string {

	// create layout based on reference date
	ref := "2006-01-02T15:04:05-07:00"
	loc, _ := time.LoadLocation("Europe/Berlin")
	t, err := time.Parse(ref, date)
	if err != nil {
		return ""
	}
	return t.In(loc).Format("02.01.2006 um 15:04")
}
