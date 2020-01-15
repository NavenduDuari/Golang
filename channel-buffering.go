package main

import (
	"fmt"
)

func f1(c chan string) {
	c <- "The msg is going from f1 to f2 via main"
	c <- "This is the second msg "
	close(c)
}

func f2(c chan string) {
	msg := <-c
	fmt.Println(msg)
	// msg = <-c
	// fmt.Println(msg)

}

func main() {

	// By default channels are unbuffered,
	// meaning that they will only accept sends (chan <-)
	// if there is a corresponding receive (<- chan)
	// ready to receive the sent value.
	// Buffered channels accept a limited number of values
	// without a corresponding receiver for those values

	c1 := make(chan string, 2)
	c2 := make(chan string, 2)

	go f1(c1)
	go f2(c2)

	msg := <-c1
	c2 <- msg
	msg = <-c1
	c2 <- msg

	fmt.Scanln()
}
