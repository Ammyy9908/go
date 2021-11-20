package main

import (
	"fmt"
)

func main() {
	fmt.Print("Arrays,Slicings,Ranges\n")

	// Arrays
	var a [10]int
	a[0] = 1          // set elements
	fmt.Println(a[0]) //read element by element index

	// declare and intialize

	var b = []int{1, 2, 3, 4, 5, 6}

	fmt.Print(string(b[0]) + "\n")

	c := [...]int{1, 2}
	fmt.Print(c)

	// Slices

	var d []int // a slice like an array but length is not specefied

	fmt.Print(d)

	chars := []string{0: "a", 2: "c", 1: "b"}

	fmt.Println(chars)

	sliceB := b[1:3]
	fmt.Println(sliceB)

	sliceC := b[:3]
	fmt.Println(sliceC)

	// create a slice using make

	f := make([]byte, 5, 5)
	fmt.Print(f)

	// Opeartions on Arrays and slice

	// get the length of the array

	fmt.Println(len(a))

	// loop over an array

	for i, e := range b {
		fmt.Println(i, e)
	}

	// if you need only element

	for _, e := range b {
		fmt.Println(e)
	}

	// if you only need index

	for i := range b {
		fmt.Println(i)
	}

}
