package main

import (
	"fmt"	
)

func main() {
	bubbleSort([]int{5, 3, -1, 4, 2, 10, -14})
}

func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
				swap(numbers, j)
			}
		}
	}

	fmt.Println(numbers)
}

func swap(numbers []int, i int) {	
	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
}