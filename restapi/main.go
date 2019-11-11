package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

// Article struct
type Article struct{
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

// Articles array
type Articles []Article

func allArticle(w http.ResponseWriter, r *http.Request)  {
	articles := Articles{
		Article{Title: "first title", Desc: "article description", Content: "lorem ipsum dolar"},
	}

	json.NewEncoder(w).Encode(articles)

}

func homePage(w http.ResponseWriter, r *http.Request)  {
	tmplt := template.Must(template.ParseFiles("home.html"))

	tmplt.Execute(w, "data")
}

func aboutPage(w http.ResponseWriter, r *http.Request)  {
	tmplt := template.Must(template.ParseFiles("about.html"))

	tmplt.Execute(w, "data")
}
func handleRequest()  {

	r := mux.NewRouter()
	r.HandleFunc("/home", homePage)
	r.HandleFunc("/about", aboutPage)
	r.HandleFunc("/articles", allArticle)

	http.ListenAndServe(":8080", r)
}

func main() {
	handleRequest()
}