package main

import "fmt"

func main() {
	
	// IF
	x := 10
	if x > 5 {
		fmt.Println("yep")
	}	
	
	// FOR
	for i := 0; i < 5; i++ {
		fmt.Println("nop")
	}

	// SWITCH
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}

	// TAG LESS SWITCH
	
	switch {
	case x > 5:
		fmt.Println("x is greater than 5")
	case x < 5:
		fmt.Println("x is less than 5")
	default:
		fmt.Println("x is equal to 5")
	}

	// BREAK

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
	
	// CONTINUE

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	
	// SCAN
	
	var appleNumber int
	fmt.Println("Enter the number of apples:")
	num, err := fmt.Scan(&appleNumber) // num is the number of items successfully scanned, err is the error (nil if no error)
	
	if err == nil {
		fmt.Printf("num: %v\n", num)
		fmt.Printf("appleNumber: %v\n", appleNumber)
	} else {
		fmt.Println(num, err)
	}
	
}

