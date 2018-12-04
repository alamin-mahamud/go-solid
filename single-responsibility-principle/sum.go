package single

func Sum(shapes []IShape) (sum float64) {
	for _, shape := range shapes {
		sum += shape.Area()
	}
	return
}
