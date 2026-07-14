package area

import (
	
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle
type Rectangle struct {
	length  float64
	breadth float64
}

var value = 110.25

func NewRectangle(length, breadth float64)  Shape{
	return &Rectangle{
		length:  length,
		breadth: breadth,
	}
}

// func NewRectangleDemo (length, breadth float64) Shape {
// 	return Rectangle{
// 		length:  length,
// 		breadth: breadth,
// 	}
// }

func (r *Rectangle) Area() float64 {

	return r.length * r.breadth
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}



// Circle
type Circle struct {
	radius float64
}

func NewCircle(radius float64) Shape {
	return &Circle{
		radius: radius,
	}
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
