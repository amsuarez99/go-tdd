package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{width: 20, height: 10}
	got := rectangle.getPerimeter()
	want := 60.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{width: 20, height: 10}
	got := rectangle.getArea()
	want := 200.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
