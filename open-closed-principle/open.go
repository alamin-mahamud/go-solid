package open

import "fmt"

type Animal interface {
	Live() string
}

type Cat struct {
	name string
}

func (c *Cat) Live() string {
	return c.name
}

type BigCat struct {
	Cat
}

func (b *BigCat) Live() string {
	return b.name + " => it's big"
}

func Run() {
	c := &Cat{"kitty"}
	b := &BigCat{Cat{"BigKitty"}}
	fmt.Println(c.Live())
	fmt.Println(b.Live())
}
