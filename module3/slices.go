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

	// MAKE
	// Make creates a new array and returns a slice that points to it
	slice3 := make([]int, 3)
	fmt.Println(slice3) // [0 0 0]
	fmt.Println(len(slice3), cap(slice3)) // 3 3

	slice4 := make([]int, 3, 5)
	fmt.Println(slice4) // [0 0 0]
	fmt.Println(len(slice4), cap(slice4)) // 3 5

	// APPEND
	slice4 = append(slice4, 1, 2, 3, 4, 5)
	fmt.Println(slice4) // [0 0 0 1 2 3 4 5]
	fmt.Println(len(slice4), cap(slice4)) // 8 10

	// Append to a slice that has enough capacity to fit the new elements
	slice5 := []int{1, 2, 3, 4, 5}
	slice5 = append(slice5, 6, 7, 8, 9, 10)
	fmt.Println(slice5) // [1 2 3 4 5 6 7 8 9 10]
	

}
