package single

// Rectangle Type should inherit IShape
type Rectangle struct {
	length, width float64
}

// Area will be used to Get the Rectangle Area
func (r *Rectangle) Area() float64 {
	return r.length * r.width
}
