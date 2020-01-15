package main

import (
	"fmt"
)

func add(a int, b int) int {
	return (a + b)
}

func add3(a, b, c int) int {
	return (a + b + c)
}

func main() {

	res := add(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = add3(1, 3, 5)
	fmt.Println("1 + 3 + 5 = ", res)
}
