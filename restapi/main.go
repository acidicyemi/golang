package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

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

	http.ListenAndServe(":8080", r)
}

func main() {
	handleRequest()
}