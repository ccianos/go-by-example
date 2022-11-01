package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
	shape() string
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (r rect) shape() string {
	s := fmt.Sprintf("\u25AD Rectangle with a width of %f and height of %f", r.width, r.height)
	return s
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) shape() string {
	s := fmt.Sprintf("\u25EF Circle with radius of %f", c.radius)
	return s
}

func measure(g geometry) {
	fmt.Println(g.shape())
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
	fmt.Println()
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
