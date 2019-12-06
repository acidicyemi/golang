package main

import (
	"fmt"
)

func main() {
	var apple int
	// var orange int32

	// invalid operation: apple == orange (mismatched types int and int32)
	// if apple == orange {
	// 	fmt.Println("they are not the same")
	// }

	// get the type of apple

	apple = 104

	getValue := string(apple)
	fmt.Println(getValue)

	// you can only convert int type to a string not even int32
	// you can also convert a byte of int into string
	fmt.Println(string([]byte{104, 105, 55}))
	fmt.Println(10.5/2)

}