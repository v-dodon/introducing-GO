package chapter_7

import "math"

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r *Rectangle) area() float64 {
	return r.height * r.width
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}