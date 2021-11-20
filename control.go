package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Control Structures")

	var isRainy bool = !true

	if isRainy {
		fmt.Println("Today is Rainy")
	} else {
		fmt.Println("Tody isn't Rainy")
	}

	// loops

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// switch

	os := "darwin"
	switch os {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Windows")
	}

	switch OS := runtime.GOOS; OS {
	case "darwin":
		fmt.Print("Mac Os")
	default:
		fmt.Print("Other Os")
	}

}
