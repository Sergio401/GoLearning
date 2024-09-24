package main

import "fmt"

func main() {
	// Hash Table
	// Hash Table is a data structure that stores key-value pairs
	// It uses a hash function to compute an index into an array of buckets or slots
	// from which the desired value can be found

	// Hash Function
	// Hash function takes a key and returns an index into the array
	// It should be fast and deterministic
	// It should distribute keys uniformly across the array

	// Hash Table in Go
	// Hash Table is implemented using a slice of empty interfaces
	// The hash function is the built-in len() function
	// The index is the return value of the hash function
	// The value is the return value of the hash function

	hashTable := make(map[string]int)
	hashTable["John"] = 20
	hashTable["Doe"] = 25
	fmt.Println(hashTable) // map[John:20 Doe:25]

	
	// Map Literal
	mapExample := map[string]int{
		"Sergio": 28,
		"Sebastian": 19,
	}
	fmt.Println(mapExample) // map[Sergio:28 Sebastian:19]

	// Accessing a value
	fmt.Println(mapExample["Sebastian"]) // 19
	
	value, p := mapExample["Sebastian"]
	fmt.Println(value, p) // 19 true

	// Iterating over a map
	for key, value := range mapExample {
		fmt.Println(key, value)
	}
}
