package shapes

func Perimeter(width, height float64) (perimeter float64) {
	if width < 0 || height < 0 {
		perimeter = 0.0
	} else {
		perimeter = 2 * (width + height)
	}
	return
}
