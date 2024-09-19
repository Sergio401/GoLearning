package main

import (
	"fmt"
	"math"
)


func main() {
	var floatNumber float64

	fmt.Println("Enter a float number:")
	_, err := fmt.Scan(&floatNumber)

	if err == nil {
		truncatedNumber := math.Trunc(floatNumber)
		intNumber := int(truncatedNumber)
		fmt.Println("Truncated number:", intNumber)
	} else {
		fmt.Println("error scanning float number:", err)
	}
}