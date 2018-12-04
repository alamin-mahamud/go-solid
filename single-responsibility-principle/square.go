package single

// Square is like rectangle but it's arms are equal
type Square struct {
	arm float64
}

// Area Implement IShape Method
func (s *Square) Area() float64 {
	return s.arm * s.arm
}
