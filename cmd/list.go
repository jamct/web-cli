package cmd

import (
	"fmt"
)

func list() {

	l := newsList(5)

	for i := 0; i < len(l.News); i++ {
		item := l.News[i]
		fmt.Println(item.Title + ": " + item.ID)
	}
}