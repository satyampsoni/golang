package main

import "fmt"

func main() {
	// switch case statements in go are similar to other languages like C, C++, Java, etc.
	// the only difference is that we don't need to put the condition in parenthesis
	// and we don't need to put curly braces around the condition
	// but we need to put curly braces around the body of the switch statement

	// syntax of switch case statement is : switch condition { case condition: body case condition: body default: body }

	// write  a switch case statement to print the day of the week

	var day string
	fmt.Print("Enter the day of the week: ")
	fmt.Scan(&day)

	switch day {
	case "Monday":
		fmt.Println("It's Monday.")
	case "Tuesday":
		fmt.Println("It's Tuesday.")
	case "Wednesday":
		fmt.Println("It's Wednesday.")
	default:
		fmt.Println("It's some other day.")
	}
}
