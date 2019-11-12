package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func handleRequest()  {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/users", allUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func allUser(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "all user endpoint")
}

func main() {
	handleRequest()
}