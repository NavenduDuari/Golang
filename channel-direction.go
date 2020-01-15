package main

import "fmt"

func ping(pingCh chan<- string) {
	pingCh <- "Msg from ping via pong to main"
}

func pong(pingCh <-chan string, pongCh chan<- string) {
	msg := <-pingCh
	pongCh <- msg

}
func main() {

	// When using channels as function parameters,
	// you can specify if a channel is meant to only
	// send or receive values. This specificity increases
	// the type-safety of the program.

	pingCh := make(chan string, 1)
	pongCh := make(chan string, 1)
	ping(pingCh)
	pong(pingCh, pongCh)

	fmt.Println(<-pongCh)
}
