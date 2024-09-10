package main

import "fmt"

func main() {
    str := "Hello"
    num := 123
    boolean := true

    fmt.Printf("String: %s\n", str)  // %s para strings
    fmt.Printf("Entero: %d\n", num)  // %d para enteros
    fmt.Printf("Booleano: %t\n", boolean) // %t para booleanos
    fmt.Printf("Hexadecimal: %x\n", num) // %x para enteros en hexadecimal
    fmt.Printf("Tipo: %T\n", num)    // %T para mostrar el tipo de variable

	// Type conversions

	numInString := float32(num)
	fmt.Printf("Valor: %v\n", numInString)
	fmt.Printf("Tipo: %T\n", numInString)

	impedance := complex(2, 3)
	fmt.Println(impedance)
}