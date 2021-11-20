package main

import (
	"fmt"
)

func main() {
	fmt.Print("Maps\n")
	var m map[string]int
	m["key"] = 42
	fmt.Println(m["key"])

}
