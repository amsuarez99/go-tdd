package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{width: 20, height: 10}
	got := rectangle.Perimeter()
	want := 60.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func checkArea(t *testing.T, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: &Rectangle{width: 12, height: 6}, want: 72.0},
		{name: "Circle", shape: &Circle{radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: &Triangle{base: 12, height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := &Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := &Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
