package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.base * t.height)
}

func Perimeter(width, height float64) (perimeter float64) {
	if width < 0 || height < 0 {
		perimeter = 0.0
	} else {
		perimeter = 2 * (width + height)
	}
	return
}
