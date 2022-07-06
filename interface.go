package main

import (
	"fmt"
	"math"
)

// interfaces are named collections of method signatures

type geometry interface {
	area() float64
	perim() float64
}

type rect2 struct { // 'rect' can't be redeclared in this package
	width, height float64
}

type circle struct {
	redius float64
}

func (r rect2) area() float64 {
	return r.width * r.height
}

func (r rect2) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.redius, 2)
}

func (c *circle) perim() float64 {
	return 2 * math.Pi
}

// If a variable has an interface type, then we can call methods that are in the named interface.
// Here’s a generic measure function taking advantage of this to work on any geometry.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	r := rect2{width: 3, height: 4}
	c := circle{redius: 5}

	// The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
	measure(r)
	measure(&c)
}

// @refer https://jordanorelli.tumblr.com/post/32665860244/how-to-use-interfaces-in-go
