// for loop in go

package main

import "fmt"

func main() {
	// for loop in go
	// for loop is used to iterate over a block of code
	// syntax of for loop in go

	// for initialization; condition; increment/decrement {
	// 	// code to be executed
	// }

	// printing numbers 10 to 1

	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	// short hand for loop
	// syntax of short hand for loop in go

	// for condition { // code to be executed } // no initialization and increment/decrement in short hand for loop

	// printing numbers 1 to 10

	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	// syntax of infinite loop in go

	// for { // code to be executed }
	// example of infinite loop

	for {
		fmt.Println("Infinite loop")
	}

}
