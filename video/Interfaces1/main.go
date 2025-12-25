package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {

	fmt.Println("Try programiz")

	r := rect{width: 5, height: 10}
	c := circle{radius: 7}
	var s shape = c
	fmt.Printf("Shape: %T\n", s)
	fmt.Printf("Area: %.2f\n", s.area())
	fmt.Printf("Perimeter: %.2f\n", s.perimeter())

	s = r
	fmt.Printf("Shape: %T\n", s)
	fmt.Printf("Area: %.2f\n", s.area())
	fmt.Printf("Perimeter: %.2f\n", s.perimeter())

}
