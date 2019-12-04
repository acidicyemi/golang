package main

import (
	"fmt"
)

func main() {

	// how to declare a channel best pratise
	c := make(chan int)
	
	// alternative 
	// var c chan int
	
	// channel receiving
	// friends := "no new friends" 
	c<- 4

	fmt.Printf("%+v\n%T\n", c, c)
}