package single

import "fmt"

// Run is the Driver of this Package
func Run() {
	s := []IShape{}
	c := Circle{10.00}
	r := Rectangle{10.00, 20.00}
	sq := Square{10.00}
	s = append(s, &c)
	s = append(s, &r)
	s = append(s, &sq)
	fmt.Println(Sum(s))
}
