// an interface is a type that defines a set of method signatures, not their implementation

package main

import (
	"fmt"
	"math"
	// "math"
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
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height // 6+8 = 14
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

func deleteCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("Circle with radius:", c.radius)
	}
}

func main() {

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	deleteCircle(r)
	deleteCircle(c)

}
