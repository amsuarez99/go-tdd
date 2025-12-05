package structs

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Shape interface {
	Area() float64
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r *Rectangle) Area() float64 {
	return (r.width * r.height)
}

type Circle struct {
	radius float64
}

func (c *Circle) Perimeter() float64 {
	return math.Pi * c.radius
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	base   float64
	height float64
}

func (c *Triangle) Area() float64 {
	return (c.base * c.height) / 2
}
