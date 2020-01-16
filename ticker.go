package main

import (
	"fmt"
	"time"
)

func main() {

	// Timers are for when you want to do something once in the future
	// - tickers are for when you want to do something repeatedly
	// at regular intervals. Here’s an example of a ticker that
	// ticks periodically until we stop it.

	ticker1 := time.NewTicker(time.Millisecond * 500)

	done := make(chan bool)

	go func() {

		for {
			select {
			case <-done:
				break
			case <-ticker1.C:
				fmt.Println("Ticker 1 fired")
			}
		}
	}()

	time.Sleep(time.Second * 3)
	ticker1.Stop()
	done <- true
	fmt.Println("Ticker 1 stopped")
}
