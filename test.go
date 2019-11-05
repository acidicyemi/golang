package main

import (
	"time"
	"fmt"
)


func main() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("this is gonna run first")
}

func hello()  {
	fmt.Println("this is gonna run when?")
}