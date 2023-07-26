package main

type Rectangle struct {
	width  int
	height int
}

func (Rec Rectangle) calculateArea() int {
	return Rec.width * Rec.height
}
