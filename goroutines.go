package main

import (  
    "fmt"
    "time"
)

func numbers() {  
    for i := 1; i <= 26; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}
func alphabets() {  
    for i := 'a'; i <= 'z'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}
func main() {  
	go alphabets()
    go numbers()
    time.Sleep(11 * time.Second)
    fmt.Println("main terminated")
}