package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}
func (r rect) perim() float64 {
	return r.height*2 + r.width*2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 10, height: 30}
	c := circle{radius: 2.5}

	// Assigning b to an interface value of type Stringer sets both words of the
	// interface value
	var b rect
	b = rect{width: 1, height: 1}
	measure(b)

	measure(r)
	measure(c)
}
