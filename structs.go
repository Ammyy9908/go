package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")
	// a struct is a type.It's also known as a collection of fields

	// Declration

	type Vertex struct {
		X, Y int
	}

	// Creating

	var v = Vertex{1, 2}
	// create a struct using keys to values

	var x = Vertex{X: 1, Y: 2}

	// accessing

	fmt.Println(x.X, x.Y)
	fmt.Println(v.X, v.Y)

}
