package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "Message 1"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("Timeout 1")
	}

	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "Message 2"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("Timeout 2")
	}

}
