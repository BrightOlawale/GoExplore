package main

import "fmt"

// AuthProvider We can use interface for creating authentication services for different implementations
type AuthProvider interface {
	Authenticate(username string, password string) (bool, error)
}

// Now, let's create 2 different implementations of AuthProvider interface

// DefaultAuthProvider A default authentication services
type DefaultAuthProvider struct{}

// ExternalAuthProvider An external authentication
type ExternalAuthProvider struct{}

// Authenticate Implementation of DefaultAuthProvider interface
func (d *DefaultAuthProvider) Authenticate(username string, password string) (bool, error) {
	if username == "admin" && password == "1234" {
		return true, nil
	}
	return false, nil
}

// Authenticate Implementation of ExternalAuthProvider interface
func (e *ExternalAuthProvider) Authenticate(username string, password string) (bool, error) {
	if username == "user" && password == "0000" {
		return true, nil
	}
	return false, nil
}

// AuthenticateUser Authentication function
func AuthenticateUser(authProvider AuthProvider, username string, password string) (bool, error) {
	return authProvider.Authenticate(username, password)
}

// Test the above
func testMoreInterfaces() {
	defaultAuthProvider := &DefaultAuthProvider{}
	externalAuthProvider := &ExternalAuthProvider{}

	// Assign values
	username, password := "admin", "1234"

	// Let's test the DefaultAuthProvider
	defaultAuthResult, _ := AuthenticateUser(defaultAuthProvider, username, password)

	// Print result
	fmt.Println("Default authentication result: ", defaultAuthResult)

	username, password = "user", "0000"
	// Let's test the ExternalAuthProvider
	externalAuthResult, _ := AuthenticateUser(externalAuthProvider, username, password)

	// Print result
	fmt.Println("External authentication result: ", externalAuthResult)
}
