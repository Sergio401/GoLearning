package main

import "fmt"

func main() {
	// SLICE

	arr := [...]string{"a", "b", "c", "d", "e", "f"}
	slice := arr[2:4] // 2 is inclusive and 4 is exclusive
	fmt.Println(slice) // [c d]
	fmt.Println(len(slice), cap(slice)) // 2 4

	// SLICE LITERAL

	// Important:Slice literal create a new array and create a slice that points to all array
	slice2 := []int{1, 2, 3, 4, 5} 
	fmt.Println(slice2) // [1 2 3 4 5]
}