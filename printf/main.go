package main

import (
	"fmt"
)

func main() {
	// %d - means an integer value
	var (
		planet = "venus"
		distance = 261
		oribital = 224.701
		hasLife = false
	)
	fmt.Printf("%+v\n", planet)
	fmt.Printf("%+v\n", distance)
	fmt.Printf("%+v\n", oribital)
	fmt.Printf("%+v\n", hasLife)
}