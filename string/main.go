package main

import (
	"unicode/utf8"
	"fmt"
)

func main() {
	name := "ajayi adeyemi"
	fmt.Println(name)
	fmt.Println(len(name))
	// best way to determin the lenght of a string
	fmt.Println(utf8.RuneCountInString(name))
}