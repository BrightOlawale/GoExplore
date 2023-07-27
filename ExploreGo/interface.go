package main

// Shape Create shape interface to store function for different shapes
type Shape interface {
	calcArea() float64
}

// Circle Create circle type to inherit from Shape interface
type Circle struct {
	radius float64
}

// Rect Create rectangle type to inherit from Shape interface
// Note that "Rectangle" is not used here because it has been already declared in this package
type Rect struct {
	width  float64
	height float64
}

// Square Create square type to inherit from Shape interface
type Square struct {
	sides float64
}

// Now let's create the calcArea function for each shape

// calcArea for circles
func (cir Circle) calcArea() float64 {
	return 3.14 * cir.radius * cir.radius
}

// calcArea for Rectangles
func (rec Rect) calcArea() float64 {
	return rec.width * rec.height
}

// calcArea for Squares
func (sqr Square) calcArea() float64 {
	return sqr.sides * sqr.sides
}
