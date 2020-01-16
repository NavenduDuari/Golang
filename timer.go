package main

import (
	"fmt"
	"time"
)

func main() {

	//Timers represent a single event in the future.
	// You tell the timer how long you want to wait,
	// and it provides a channel that will be notified at
	// that time. This timer will wait 2 seconds.

	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second * 3)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	isStop := timer2.Stop() //bool func , returns true if timer is stopped
	if isStop {
		fmt.Println("TImer 2 stopped")
	}
}
