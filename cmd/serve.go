package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func serve(){
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/news", handleAll)
	r.HandleFunc("/news/{id}", handleSingle)

	r.HandleFunc("/api/news", handleAllApi).Methods("GET")
	r.HandleFunc("/api/news/{id}", handleSingleApi).Methods("GET")

	go log.Fatal(http.ListenAndServe(":8080", r))
}

func handleAll(w http.ResponseWriter, r *http.Request) {
	 fmt.Fprint(w, "<h1>News von heise.de</h1>")
	 news := newsList(5)
	 for i := 0; i < len(news.News); i++ {
		item := news.News[i]
		fmt.Fprintf(w, `
		<hr/>
		<h2><a href='news/`+item.ID+`'>%v</a></h2>
		<h3>Veröffentlicht am %v</h3>
		<p>%v</p>
		`, item.Title, formatDate(item.Date), item.Description )
		
	}	 
}

func handleSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	news, err := newsSingle(id)
	if err != nil{
		w.Write([]byte("not found"))
		return
	}
	fmt.Fprintf(w, `
		<h1>%v</h1>
		<h3>Veröffentlicht am %v</h3>
		<p>%v %v</p>
		`, news.Title, formatDate(news.Date), news.Description, news.ID )
}

func handleAllApi(w http.ResponseWriter, r *http.Request) {
	news := newsList(5)
	js, _ := json.Marshal(news)
	w.Write(js)
}


func handleSingleApi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	s, err := newsSingle(id)
	if err != nil{
		w.Write([]byte(`{"error": "not found"}`))
		return
	}
	js, _ := json.Marshal(s)
	w.Write(js)
}