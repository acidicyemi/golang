package main

import (
	// "fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email string
	Subject string
	Message string
}

func main() {

	templt := template.Must(template.ParseFiles("forms.html"))
	// fmt.Println(templt)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
            templt.Execute(w, nil)
            return
        }

        details := ContactDetails{
            Email:   r.FormValue("email"),
            Subject: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }

        // do something with details
        _ = details

        templt.Execute(w, struct{ Success bool }{true})
	})
	http.ListenAndServe(":8080", nil)
}