package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	wight  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.wight
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perimeter() float64 {
	return 2 * (r.wight + r.height)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rect{wight: 10, height: 20}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
