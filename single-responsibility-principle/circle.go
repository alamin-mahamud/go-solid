package single

import "math"

// Circle is a basic Shape which only has radius
type Circle struct {
	radius float64
}

// Area is the implemented method for Circle Shape. That's how it is implementing IShape.
func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
