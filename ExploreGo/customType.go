package main

// User :Creating a custom type
type User struct {
	Id   int
	Name string
}

// Note that the custom type is exported because it starts with a capital letter
// If it starts with a lowercase letter, it is not exported
