package main

import (
	"fmt"
	"log"
	"net/http"
)

// Books Creating Book type
type Books struct {
	Id    int
	Title string
}

// Creating Book data
var allBooks []Books = []Books{
	{
		Id:    1,
		Title: "The Go Programming Language",
	},
	{
		Id:    2,
		Title: "Python: A definitive guide",
	},
	{
		Id:    3,
		Title: "Java: Understanding OOP in Java",
	},
}

// The main function: Entry point of the program
func main() {
	// First, we create a mux to handle our routes.

	// A mux is an HTTP request multiplexer that matches the URL of incoming requests
	// against a list of registered patterns and calls an associated handler for
	//the pattern that most closely matches the URL

	var mux *http.ServeMux = http.NewServeMux()
	// Note that "*http.ServeMux" is a pointer to a struct of type "http.ServeMux"
	// that implements the "http.Handler" interface

	// Now, let's register our routes with the mux we created
	mux.HandleFunc("/books", getAllBooks)
	mux.HandleFunc("/books/create", createBook)

	// Now, let's create our server with the mux we created as the handler and
	// 8080 as the port number (address) of the server
	var server *http.Server = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Now, let's start the server
	fmt.Println("Starting server...")

	// Check if an error occurred while starting the server
	if err := server.ListenAndServe(); err != nil {
		// panic if we failed to connect
		log.Panic(err)
	}

}
