package main

import "fmt"

// Function to double the value of a number
func double(num *int) {
	*num *= 2
}

// Just to understand pointers
func pointers() {
	var a int = 10
	// or
	// a := 10

	// Pointer to num
	var ptr *int = &a
	// or
	// ptr := &a

	// Checking value of a
	fmt.Println("Value of a:", a)
	fmt.Println("Actual address of a:", ptr)
	fmt.Println("Value of a using pointer:", *ptr)

	// Changing value of a using pointer
	*ptr = 20
	fmt.Println("Value of a:", a)
	fmt.Println("Actual address of a:", ptr)
	fmt.Println("Value of a using pointer:", *ptr)
}

// Function to a slice as an argument
func workSlice(slice []int) {
	// Changing the value of the first element
	slice[0] = 25
}

// Main function
func main() {
	pointers()

	// Creating a variable
	num := 30

	// Double the value of num
	double(&num)

	// Print the value of num
	fmt.Println("Doubled of num:", num)

	// Create a slice of integers
	slice := []int{40, 50, 60, 70, 80}

	// Print the slice after passing it to the function
	fmt.Println("Before passing to func:", slice)
	workSlice(slice)
	fmt.Println("After passing to func:", slice)

	// Function to test the custom type
	customTypeTest()

	// Function to test the hashPassword method
	hashTest()
}
