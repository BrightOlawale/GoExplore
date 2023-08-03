package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

// Learning error handling
func division(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by 0")
	}
	return i / j, nil
}

// CustomError Custom errors can be created like this:
type CustomError struct {
	message string
}

// Attaching the custom error to the error interface
func (er CustomError) Error() string {
	return er.message
}

// errors.Wrap allows us to add context to errors
func cause() error {
	return errors.New("cause Error")
}

// Now using the wrap method
func processTwo() error {
	return errors.Wrap(cause(), "error occurred while processing data")
}

// Creating a function to test my custom error
func processError() error {
	return CustomError{message: "my custom error"}
}

// Handling Errors with multiple return values
func readFile(filePath string) ([]byte, error) {
	// Open the file
	file, err := os.Open(filePath)

	// If error occurred, return nil and the error
	if err != nil {
		return nil, err
	}

	// Defer closing the file after the function returns
	defer file.Close()

	// Read the file
	data := make([]byte, 1024)
	count, err := file.Read(data)

	// If error occurred, return nil and the error
	if err != nil {
		return nil, err
	}

	// Return the data and nil
	return data[:count], nil
}

// Testing error handling
func testError() {
	// Test the division function
	div, err := division(6, 0)
	if err != nil {
		fmt.Println("Error (division by 0): ", err)
	} else {
		fmt.Println("Result: ", div)
	}

	// Testing custom error
	errr := processError()

	if errr != nil {
		fmt.Println("ERROR: ", errr)
	}

	// Testing another custom error
	processErr := processTwo()

	if processErr != nil {
		fmt.Println("ERROR:", processErr)
	}

	// Testing error handling with multiple return values
	data, err := readFile("test.txt")
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println("Data: ", data)
	}
}
