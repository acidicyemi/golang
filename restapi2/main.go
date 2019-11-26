package main

import (
	"fmt"
	"runtime"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func handleRequest()  {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", Homepage).Methods("GET")
	r.HandleFunc("/users", AllUser).Methods("GET")
	r.HandleFunc("/newuser/{name}/{email}", NewUser).Methods("Post")
	r.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/user/{name}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	
	numberOfCpu := runtime.NumCPU()
	fmt.Println( numberOfCpu)
	handleRequest()
}