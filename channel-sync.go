package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second * 2)
	fmt.Println("done")
	done <- true
	close(done)
}
func main() {

	done := make(chan bool)
	go worker(done)

	fmt.Println("working done:", <-done)
}
