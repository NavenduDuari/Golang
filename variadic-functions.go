package main

import (
	"fmt"
)

func sum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}

	return total
}
func main() {
	fmt.Println("sum(1, 2, 3) =", sum(1, 2, 3))
	nums := []int{1, 2, 3, 4}
	fmt.Println("sum of nums =", sum(nums...))
}
