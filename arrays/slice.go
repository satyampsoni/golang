package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	slice := arr[1:3]
	fmt.Println(slice)
}
