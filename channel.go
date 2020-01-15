package main

import (
	"fmt"
)

func f1(c chan string) {
	c <- "The msg is going from f1 to f2 via main"
	close(c)
}

func f2(c chan string) {
	msg := <-c
	fmt.Println(msg)
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go f1(c1)
	go f2(c2)

	msg := <-c1
	c2 <- msg

	fmt.Scanln()
}
