package main

import (
	"fmt"
)

func vals() (int, int) {
	return 1, 3
}

func main() {

	v1, v2 := vals()
	fmt.Println("v1 = ", v1, "v2 = ", v2)

	_, a := vals()
	fmt.Println("a = ", a)
}
