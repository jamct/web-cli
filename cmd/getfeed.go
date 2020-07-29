package cmd

import (
	"errors"
	"strings"

	"github.com/mmcdole/gofeed"
)

type NewsList struct{
	News []News	`json:"news,omitempty"`
	BaseUrl string `json:"base_url,omitempty"`
}
type News struct{
	Date		string	`json:"date,omitempty"`
	ID 			string		`json:"id"`
	Title		string  `json:"title"`
	Description	string	`json:"description"` 	 
}

func newsList(count int) NewsList{
	list := NewsList{ BaseUrl: "http://heise.de/"}
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")

	for i := 0; i < count; i++ {
	  item := feed.Items[i]
	  id := strings.Replace(item.GUID, list.BaseUrl, "", -1)
	  n :=	News{ item.Published, id, item.Title, item.Description}
	  list.News = append(list.News, n)
	}
	return list
}

func newsSingle(id string) (News, error){

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")

	for _, item := range feed.Items {
		if item.GUID == "http://heise.de/" +id {
			id := strings.Replace(item.GUID, "http://heise.de/", "", -1)
			  single :=	News{ item.Published, id, item.Title, item.Description}
			return single, nil
		}
	}
	return News{}, errors.New("id not found")
}