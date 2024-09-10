package main

import "fmt"

func main() {
	// most basic declaration
	var x int = 2 // keyword - name - type || default value = 0
	fmt.Println(x)
	
	var xa = 3
	fmt.Println(xa)

	/* I can use:
	 	uint8 uint16 uint32 uint64 uint
		int8 int16 int32 int64 int
		float16 float32 float
		complex16 complex32
		string
		bool
		byte: uint alias to represent a byte 
	*/


	// other ways 
	var y, z int = 2, 3
	fmt.Println(y, z)

	// This kind of declaration works only inside a function
	a := "Hello world" 
	fmt.Println(a)

	b := 2 + 2 // I can use expressions
	fmt.Println(b)

	// ---- TYPES ------
	
	type Celsius float32
	type IDnum int

	var temp Celsius = 37.4
	var pid IDnum = 1234
	fmt.Println(temp, pid)
}

