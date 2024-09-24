package main

import "fmt"

func main() {

	// Structs
	type User struct {
		ID int
		FirstName string
		LastName string
		Email string
		IsActive bool
	}
	
	// Creating a new user
	user := User{
		ID: 1,
		FirstName: "Sergio",
		LastName: "Gomez",
		Email: "sergio@gmail.com",
		IsActive: true,
	}

	fmt.Println(user)
	
	// Creating a new user
	var user2 User
	fmt.Println(user2)

	// Updating the user
	user2.ID = 2
	user2.FirstName = "Sebastian"
	user2.LastName = "Gomez"
	user2.Email = "sebastian@gmail.com"
	user2.IsActive = true

	fmt.Println(user2)
}
