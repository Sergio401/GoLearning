package main

import "fmt"

func main() {
	
	// ARRAYS
	
	var numbers [5]int
	numbers[0] = 5
	fmt.Println(numbers) // [5 0 0 0 0]

	// ARRAYS LITERAL
	var digits [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(digits) // [1 2 3 4 5]

	odds := [...]int{1, 3, 5, 7, 9} // I can use ... to let the compiler count the number of elements
	fmt.Println(odds) // [1 3 5 7 9]

	// Iterate over an array
	for i, value := range odds {
		fmt.Println("Index:", i, "Value:", value) // Index: 0 Value: 1
	}
	
	

}