package main

import "fmt"

func main() {
	// if else statements in go are similar to other languages
	// the only difference is that we don't need to put the condition in parenthesis
	// and we don't need to put curly braces around the condition
	// but we need to put curly braces around the body of the if statement
	// and the else statement
	// also we need to put the else statement in the same line as the closing curly brace of the if statement
	// otherwise we will get a syntax error

	if 1 == 1 {
		fmt.Println("1 is equal to 1")
	} else {
		fmt.Println("1 is not equal to 1")
	}

	// we can also use the else if statement
	// the else if statement is similar to the if statement
	// the only difference is that we need to put the else if statement in the same line as the closing curly brace of the if statement
	// syntax of else if statement is : if condition { body } else if condition { body } else { body }

	if 2 == 2 {
		fmt.Println("2 is equal to 2")
	} else if 2 == 1 {
		fmt.Println("2 is equal to 1")
	} else {
		fmt.Println("2 is not equal to 2")
	}

	// we can also use the if statement without the else statement
	// syntax of if statement without else statement is : if condition { body }

	if 3 == 3 {
		fmt.Println("3 is equal to 3")
	}

}
