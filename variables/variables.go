package main

import "fmt"

func main() {

	// var is used to declare a variable
	var name = "satyam soni"
	fmt.Println("Hello", name)

	// var can declare multiple variables at once
	var a, b = 1, 2
	fmt.Println(a, b)

	// const is used to declare a constant
	// const is used for those variables whose value is not going to change
	const c = "constant"
	fmt.Println(c)

	// changing the values of a and b
	a, b = 3, 4
	fmt.Println(a, b)

	// we can't change the value of c as it is a constant

	// c = "changed" - gives an error

}
