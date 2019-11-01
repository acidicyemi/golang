package main

import (
	"fmt"
	"net/http"
)

// import 
func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("this is a test file")
	})

	http.ListenAndServe(":8080", nil)
}