package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	
	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "devuser:password@(127.0.0.1:3306)/dbname?npmp=true")

	// if there is an error inserting, handle it
	if err != nil {
		// panic(err.Error())
		fmt.Println("hi")
	}
	// be careful deferring Queries if you are using transactions
	defer db.Close()
	fmt.Println("able to connect")
}