package cmd

import (
	"fmt"
	"time"
)

func describe(id string) {

	id = "-"+id
	s, err := newsSingle(id)

	if err != nil{
		fmt.Println("Article " + id + " not found.")
		return
	}

	fmt.Println(formatDate(s.Date))
	fmt.Println(s.Title)
	fmt.Println(" ")
	fmt.Println(s.Description)
	return
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