package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	// run go biuld -o greeter
	// the -o gives the program a name

	// run command 
	// ./greeter adeyemi ajayi
	// []string{"./greeter", "adeyemi", "ajayi"}

	// to return the argument individually

	ls := len(os.Args)
	for l := 0; l < ls; l++ {
		fmt.Printf("the argiment of %v is %v\n", l, os.Args[l])
	}

}