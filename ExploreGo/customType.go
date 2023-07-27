package main

import "fmt"

// User :Creating a custom type
type User struct {
	Id   int
	Name string
}

// Note that the custom type is exported because it starts with a capital letter
// If it starts with a lowercase letter, it is not exported

// Function to test the above codes
func customTypeTest() {
	// Using the custom type in main function
	// Assigning values to the custom type
	firstUser := User{
		Id:   1,
		Name: "Olawale",
	}

	// Printing the values
	fmt.Println("Id: ", firstUser.Id)
	fmt.Println("Username: ", firstUser.Name)

	// Using the calculateArea function
	// Assign values
	rectangle := Rectangle{width: 5, height: 5}

	// Print the calculated area
	fmt.Println("Area of the rectangle: ", rectangle.calculateArea())
}
