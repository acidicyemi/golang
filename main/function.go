package main

import (
	"fmt"
)

func main() {
	area, perimeter := rect(2.3,3.0)
	fmt.Println(area, perimeter )
	fmt.Printf("area %t and parameter %t", int(area), float32(perimeter))

}

func rect(length, width float64) (float64, float64) {
	area := length * width
	perimeter := (length * width) *2
	return area, perimeter
}