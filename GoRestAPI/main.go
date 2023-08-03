package main

import (
	"encoding/json"
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

// getAllBooks: return all books data from the server
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	// Convert the book data to JSON format
	response, err := json.Marshal(allBooks)

	// If error occurred
	if err != nil {
		log.Panic(err)
	}

	// Now let us set the response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // StatusOk is 200

	// Now let us set the response body
	_, err = w.Write(response)
	if err != nil {
		log.Panic(err)
	}

}

// createBook: create a new book in our record
func createBook(w http.ResponseWriter, r *http.Request) {
	// Let's check the request method, if it is not POST, we return an error
	// We can check the request method using the "r.Method" field
	// which returns the request method
	// http.MethodPost is a constant that represents the POST method
	if r.Method != http.MethodPost {
		// Then response header is set to 405 (Method Not Allowed)
		// and we use w.Write to write the error message to the response body
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method not allowed"))
		if err != nil {
			log.Panic(err)
		}
		return
	}

	// If method is not a problem, then we can proceed to create the book
	// First, we create a new book variable of type Books
	var newBook Books

	// Now, we decode the request body into the newBook variable
	// Note that we use the json.NewDecoder.Decode method to decode the request body
	// into the newBook variable
	// The first argument is the request body and the second argument is the variable
	// we want to decode the request body into
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		log.Panic(err)
	}

	// Now, we append the newBook to the allBooks slice
	allBooks = append(allBooks, newBook)

	// Now, we encode the newBook variable into JSON format
	// using the json.Marshal method
	response, err := json.Marshal(newBook)
	if err != nil {
		log.Panic(err)
	}

	// Now, we set the response header to 201 (Created)
	// and use w.Header().Set to set the content type to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Now, we write the response to the response body
	_, err = w.Write(response)
	if err != nil {
		log.Panic(err)
	}
}
