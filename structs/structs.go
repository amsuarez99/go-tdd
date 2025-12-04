package structs

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) getPerimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r *Rectangle) getArea() float64 {
	return (r.width * r.height)
}
