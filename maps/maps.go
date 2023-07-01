package main

import "fmt"

func main() {
	// A map is a collection of key-value pairs. Maps are also known as associative arrays, hash tables, or dictionaries.

	// syntax to declare a map is: map[keyType]valueType

	map1 := make(map[string]int)

	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3

	fmt.Println("Map1", map1)

	//maps are unordered we can't access them by index

}
