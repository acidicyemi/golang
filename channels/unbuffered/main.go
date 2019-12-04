package main

import (
	// "time"
	"fmt"
	"sync"
)

var waitG sync.WaitGroup

func main() {
	c := make(chan int)
	
	waitG.Add(1)
	go send(c)
	go receive(c)
	waitG.Wait()

	// time.Sleep(40 * time.Microsecond)
}

func send(c chan int)  {
	for i := 0; i < 6; i++ {
		c <- i
	}
	waitG.Done()
}

func receive(c chan int)  {

	// for {
	for i := 0; i < 2; i++ {
		fmt.Println(<-c)
	}	
	waitG.Done()
}
