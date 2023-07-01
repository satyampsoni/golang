package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array", arr)

	slice := arr[1:3]
	fmt.Println("slice of the array is", slice)

	// value from the beginning to the 3rd index
	slice1 := arr[:3]
	fmt.Println("First three numbers are", slice1)

	// value from the 1st index to the end
	slice2 := arr[1:]
	fmt.Println("Last four numbers are", slice2)

	//removing a value from the slice
	slice3 := append(arr[:2], arr[3:]...)
	fmt.Println("After removing the value", slice3)

	// the first part of the slice is the low bound and the second part is the high bound and arr[:2] takes the first two values from the array and arr[3:] takes the values from the 4th index to the end of the array and then we append the two slices together to get the final slice.
}
