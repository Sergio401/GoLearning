package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	
	var input string
	slice := make([]int, 3)

	for i := 0; i < cap(slice); i++ {
		fmt.Println("Enter a number:")
		fmt.Scan(&input)
		
		if strings.ToLower(input) == "x" {
			fmt.Println("Exiting program.")
			break
		}

		// Convert input string to int
		value, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Error reading input:", err)
		} else {
			if i < 2 {
				slice[i] = value
				sort.Ints(slice)
				fmt.Println(slice)
			}			
			if i >= 2 {
				slice = append(slice, value)
				sort.Ints(slice)
				fmt.Println(slice)
			}
		}
		
	}
}
