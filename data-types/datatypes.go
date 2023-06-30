// Data types in Go
// Three main categories of data types in Go: strings, integers and booleans.

// 1. Strings
// Strings are a sequence of characters. They are declared using double quotes.
package main

import "fmt"

func main() {
	// declaring a string
	var name string = "satyam soni"
	fmt.Println(name)

	// declaring a string without specifying the type
	var name1 = "Go programming"
	fmt.Println(name1)

	// declaring a string without using var keyword
	name2 := "Go programming"
	fmt.Println(name2)

	// declaring a multi-line string
	name4 := `Go is a Programming language
	which is used to develop web applications`
	fmt.Println(name4)

	// 2. Integers
	// Integers are whole numbers. They can be positive or negative.

	// declaring an integer
	var age int = 20
	fmt.Println(age)

	// declaring an integer without specifying the type
	var age1 = 30
	fmt.Println(age1)

	// declaring an integer without using var keyword
	age2 := 40
	fmt.Println(age2)

	// 3. Booleans
	// Booleans are either true or false.

	// declaring a boolean
	var isTrue bool = true
	fmt.Println(isTrue)

	// declaring a boolean without specifying the type
	var isTrue1 = true
	fmt.Println(isTrue1)

	// declaring a variable without using var keyword
	isFalse := false
	fmt.Println(isFalse)

}
