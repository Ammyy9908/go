package main

import (
	"fmt"
)

// primitive type
/*
bool
numeric type
*/

func main() {
	// bools
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v %T\n", n, n)
	fmt.Printf("%v %T\n", m, m)

	// integer
	age := 24
	fmt.Printf("%v,%T\n", age, age)

	// unsigned int

	var x uint16 = 12
	fmt.Printf("%v,%T\n", x, x)

	// strings

	var name string = "sumit"
	fmt.Println(name)

}
