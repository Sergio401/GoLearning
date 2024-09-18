package main

import "fmt"

func main() {
	var floatNumber float32

	fmt.Println("Enter a float number:")
	_, err := fmt.Scan(&floatNumber)

	if err == nil {
		truncatedNumber := int(floatNumber)
		fmt.Println("Truncated number:", truncatedNumber)
	} else {
		fmt.Println("error scanning float number:", err)
	}
}