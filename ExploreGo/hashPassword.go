package main

// Import bcrypt package
import "golang.org/x/crypto/bcrypt"

// NewUser Create type for users
type NewUser struct {
	Id       int64
	Username string
	Password string
}

// HashPassword Method to hash password of user
func (user *NewUser) HashPassword() error {
	// Notice in the above we used a pointer receiver which we did not use in the calculateArea method
	// This is because we want to modify the value of the password field of the user struct
	// If we do not use a pointer receiver, the value of the password field will not be modified
	// So use a pointer receiver when you want to modify the value of a field in a struct

	// Hash password using bcrypt algorithm
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// Note that the second argument is the cost of hashing the password, which is set to 10 by default as used above
	// cost determines the time taken to hash the password, that is the computational power required to hash the password
	// or say the complexity of the algorithm

	// Check if error occurred
	if err != nil {
		return err
	}

	// Set hashed password to user password
	user.Password = string(hashedPassword)

	// Return nil if no error occurred
	return nil
}

// Function to test the above codes for creating a new user and hashing the password
func hashTest() {
	// Create new user
	newUser := NewUser{
		Id:       1,
		Username: "test",
		Password: "WhatEverYouLike",
	}

	// Hash password
	err := newUser.HashPassword()

	// Check if error occurred
	if err != nil {
		// Print error
		println("Error occurred: ", err)
	}

	// Print password
	println("Hashed Password: ", newUser.Password)
}
