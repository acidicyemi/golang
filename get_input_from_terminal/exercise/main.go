package main

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args
	// count the length of argumnts
	fmt.Printf("there are %v argument(s) in the command\n", len(args))

	// prnt the path
	fmt.Printf("the path of the variable is %v\n", args[0])

	for arg := 1; arg < len(args); arg++ {
		fmt.Printf("hello %v\n", args[arg])
	}
}