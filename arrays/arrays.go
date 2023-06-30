package main

import "fmt"

func main() {
	// Arrays in go
	// Arrays are a collection of elements of the same data type.
	// declaring an array with type

	// int array
	var arr = [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	// string array
	var lang = [3]string{"Go", "Python", "Java"}
	fmt.Println(lang)

	// array with the size
	fmt.Println("Length of array is", len(arr))

	// another way to declare an array
	var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("size of arr1 is", len(arr1))

}
