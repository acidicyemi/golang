package main

import (
	"fmt"
)

func main() {
	ait := make(chan bool)
	go hello(ait)
	value := <- ait
	fmt.Println(value)
	fmt.Println("hello from main") 
}
func hello(ait chan bool)  {
	fmt.Println("hello from channel hello")
	ait <- true
}