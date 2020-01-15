package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.height * r.width
}

func (rPtr *rect) perimeter() int {
	return 2 * (rPtr.height + rPtr.width)
}
func main() {

	r := rect{width: 10, height: 5}
	rPtr := &r

	fmt.Println("Area of rectangle :", r.area())
	fmt.Println("Perimeter of r:", rPtr.perimeter())
}
