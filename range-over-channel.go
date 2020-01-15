package main

import "fmt"

func main() {

	queue := make(chan string, 2)

	queue <- "one"
	queue <- "two"
	close(queue) // if channel is not closed, the range func
	// wait for values and leads to deadlock

	for val := range queue {
		fmt.Println(val)
	}

}
