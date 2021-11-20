package main

import (
	"fmt"
)

// function witout any param

func sayHello() {
	fmt.Println("Hi")
}

// function with a paramter

func greetUser(name string) {
	fmt.Println("Hello " + name)
}

// function with parameter

func joinName(firstName string, lastName string) string {

	return firstName + " " + lastName
}

// function with multiple returns

func formatDetails(name string, age int) (int, string) {
	return age, name
}

func main() {
	sayHello()
	greetUser("Sumit")
	fmt.Println(joinName("sumit", "kumar"))
	fmt.Println(formatDetails("sumit", 24))

	// function as values and clousres
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(12, 13))

}
