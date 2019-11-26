package main

import (
	"fmt"
)

func main() {
	days := make(map[string]int)
	
	if days == nil {
		fmt.Println("this is an empty map")
	}else {
		fmt.Println(days)
	}

	months := map[string]int{
		"november" : 11,
		"december" : 12,
	}
	fmt.Println(months)
}