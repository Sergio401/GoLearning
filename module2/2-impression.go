package main

import "fmt"

func main() {
    str := "Hello"
    num := 123
    boolean := true

    fmt.Printf("String: %s\n", str)  // %s to strings
    fmt.Printf("Int: %d\n", num)  // %d to integers
    fmt.Printf("Boolean: %t\n", boolean) // %t to booleans
    fmt.Printf("Hexadecimal: %x\n", num) // %x to integers in hexadecimal
    fmt.Printf("Type: %T\n", num)    // %T to show the type of a variable

	// Type conversions

	numInString := float32(num)
	fmt.Printf("Value: %v\n", numInString)
	fmt.Printf("Type: %T\n", numInString)

	impedance := complex(2, 3)
	fmt.Printf("Value: %v\n", impedance)
}
