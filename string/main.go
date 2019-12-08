package main

import (
	"strings"
	"os"
	"unicode/utf8"
	"fmt"
)

func main() {
	// name := "ajayi adeyemi"
	// fmt.Println(name)
	// fmt.Println(len(name))
	// // best way to determin the lenght of a string
	// fmt.Println(utf8.RuneCountInString(name))

	word := os.Args[1]
	l := utf8.RuneCountInString(word)
	repeat := strings.Repeat("!", l )
	output := repeat+ word + repeat
	output = strings.ToUpper(output)
	fmt.Println(output)
}