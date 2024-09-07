package main

import "fmt"

// Always, the first program :)
func main() {
	var x int = 1
	var y int
	var ip *int

	ip = &x 	// ip now points to x
	y = *ip		// y is now 1

	fmt.Println(x, y, ip)
	
	// ---- NEW -----

	ptr := new(int)
	*ptr = 2

	fmt.Println(*ptr, ptr) // data and address
}

