package main

import "fmt"

type ShapeCalculator interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) Perimeter() int {
	return r.Width * r.Height
}

func (r *Rectangle) Area() int {
	return 2 * (r.Width + r.Height)
}

// both of following line guarantee that Rectangle implements all methods of Interface
// var _ ShapeCalculator = &Rectangle{}
var _ ShapeCalculator = (*Rectangle)(nil)

// newShape returns the object of interface i.e ShapeCalculator
func newShape(width int, height int) ShapeCalculator {
	return &Rectangle{
		Width:  width,
		Height: height,
	}
}

func main() {
	rect := newShape(3, 4)
	fmt.Println(rect.Area())
}
