// user input in go

package main

import "fmt"

func main() {

	// User input
	var input string

	fmt.Print("Enter your name: ")
	fmt.Scan(&input)
	fmt.Println("Hello", input)

	// User input with multiple values

	var input1, input2 string

	fmt.Print("Enter your first name: ")
	fmt.Scan(&input1)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&input2)

	fmt.Println("Hello", input1, input2)

	// Scan is used to take input from the user and & is used to store the input in the variable
	// & is a pointer which stores the address of the variable lets see an example

	fmt.Println(&input1) // 0xc000010240 this is the address of the variable input1
	fmt.Println(&input2) // 0xc000010260 this is the address of the variable input2
	fmt.Println(&input)  // 0xc000010280 this is the address of the variable input

}
