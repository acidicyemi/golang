package main

import (
	"fmt"
)

func main() {
	// Pointer
	// the location which a value is stored
	// mean a loaction which a variable is stored
	firstPointer()
}

func firstPointer() {
	x := 2
	p := &x
	q := *p + *p

	fmt.Printf("x = %T and x = %d \n", x, x)
	fmt.Printf("x = %T and x = %x \n", p, p)
	fmt.Printf("p = %T and p = %d \n", *p, *p)
	fmt.Printf("q = %T and q = %d \n", q, q)
}
