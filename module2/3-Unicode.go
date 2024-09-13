package main

import "fmt"
import "unicode"
import "strings"
import "strconv"

func main() {
    str := 'H'  // 'H' is a rune, to different from string we use '' and not ""
    
    fmt.Println(unicode.IsSpace(str)) // false
    fmt.Println(unicode.IsDigit(str)) // false
    fmt.Println(unicode.IsLetter(str)) // true
    fmt.Println(unicode.IsLower(str)) // false
    fmt.Println(unicode.IsNumber(str)) // false
    fmt.Println(unicode.IsUpper(str)) // true

    // STRING MANIPULATION

    fmt.Println(strings.Compare("Hello", "Hello")) // 0
    fmt.Println(strings.Compare("Hello", "Hell")) // 1
    fmt.Println(strings.Compare("Hello", "Helloo")) // -1

    fmt.Println(strings.Contains("Hello", "He")) // true
    fmt.Println(strings.Contains("Hello", "H")) // true
    fmt.Println(strings.Contains("Hello", "c")) // false
    fmt.Println(strings.Contains("Hello", "o")) // true
    
    fmt.Println(strings.Replace("Hello", "e", "a", 1)) // Hallo, (string, old, new, n)
    fmt.Println(strings.Replace("Hello", "l", "a", -1)) // Heaao, (string, old, new, n), n=-1 replace all ocurrences

    fmt.Println(strconv.Atoi("123")) // 123, convert string to int, return int and error
    fmt.Println(strconv.Itoa(123)) // "123", convert int to string
}
