package main

import (
	"encoding/xml"
	// "encoding/json"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Article struct
type Article struct {
	Title   string `xml:"title"`
	Desc    string `xml:"desc"`
	Content string `xml:"content"`
}

// Articles array
type Articles []Article

func allArticle(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "first title", Desc: "article 1 description", Content: "lorem ipsum dolar"},
		Article{Title: "second title", Desc: "article 2 description", Content: "lorem ipsum dolar"},
		Article{Title: "third title", Desc: "article 3 description", Content: "lorem ipsum dolar"},
	}

	xml.NewEncoder(w).Encode(articles)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmplt := template.Must(template.ParseFiles("home.html"))

	tmplt.Execute(w, "data")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	tmplt := template.Must(template.ParseFiles("about.html"))

	tmplt.Execute(w, "data")
}
func handleRequest() {

	r := mux.NewRouter()
	r.HandleFunc("/home", homePage)
	r.HandleFunc("/about", aboutPage)
	r.HandleFunc("/articles", allArticle)

	http.ListenAndServe(":8080", r)
}

func main() {
	handleRequest()
}
