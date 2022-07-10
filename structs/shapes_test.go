package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	for _, c := range casesPerimeter() {
		t.Run(c.name, func(t *testing.T) {
			got := Perimeter(c.width, c.height)
			if got != c.perimeter {
				t.Errorf("got '%f', but want '%f', given '(%f, %f)'", got, c.perimeter, c.width, c.height)
			}
		})
	}

}

type testPerimeter struct {
	name      string
	width     float64
	height    float64
	perimeter float64
}

func casesPerimeter() (cases []testPerimeter) {
	return []testPerimeter{
		{"with 0", 3, 0, 6},
		{"with negative", -2, 5, 0},
		{"both positive", 4.2, 1.5, 11.4},
	}
}

func TestArea(t *testing.T) {
	for _, c := range casesArea() {
		t.Run(c.name, func(t *testing.T) {
			got := c.shape.Area()
			if got != c.area {
				t.Errorf("got '%f', but want '%f', given '%v'", got, c.area, c.shape)
			}
		})
	}

}

type testArea struct {
	name  string
	shape Shape
	area  float64
}

func casesArea() []testArea {
	return []testArea{
		{"rectangle", Rectangle{3.0, 2.0}, 6.0},
		{"circle", Circle{10.0}, 314.1592653589793},
		{"triangle", Triangle{12.0, 6.0}, 36.0},
	}
}
