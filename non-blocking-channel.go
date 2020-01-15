package main

import "fmt"

// Basic sends and receives on channels are blocking.
// However, we can use select with a default clause
// to implement non-blocking sends, receives, and even
// non-blocking multi-way selects.

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	//non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("Message recieved:", msg)
	default:
		fmt.Println("Message NOT received")
	}
	//non-blocking send
	msg := "This is a message"
	select {
	case messages <- msg:
		fmt.Println("Message sent")
	default:
		fmt.Println("Message NOT sent")
	}

	//We can use multiple cases above the default
	//clause to implement a multi-way non-blocking select.
	// Here we attempt non-blocking receives on both
	// messages and signals.
	select {
	case <-messages:
		fmt.Println("Message received")
	case <-signals:
		fmt.Println("Signal received")
	default:
		fmt.Println("NO activity")
	}
}
