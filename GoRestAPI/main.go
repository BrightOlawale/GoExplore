package main

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
