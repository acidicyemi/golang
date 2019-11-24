package main

import (
	"net/http"
	"fmt"
)

// Homepage wr
func Homepage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Home page")
}

// AllUser wr
func AllUser(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "all user endpoint")
}

// NewUser wr
func NewUser(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "new  user endpoint")
}

// DeleteUser wr
func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "delete user endpoint")
}

// UpdateUser wr
func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "update user endpoint")
}