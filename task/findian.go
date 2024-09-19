package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter a word or phrase (or 'exit' to quit):")
		input, err := reader.ReadString('\n')
		
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		input = strings.ToLower(input)

		if input == "exit" {
			fmt.Println("Exiting program.")
			break
		}

		if strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n") {
			fmt.Println("Found!")
			fmt.Println(" ")
		} else {
			fmt.Println("Not Found!")
			fmt.Println(" ")
		}
	}
}

