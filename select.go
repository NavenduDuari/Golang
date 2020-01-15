package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}

	}()

	go func() {
		for {
			c2 <- "every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	//blocks the msg from c1 even if it is ready to send msg
	//  waiting for the c2 to finish in the loop
	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	//select allows moving forward with the channel which is ready
	//to send msg without waiting for other channel to be ready.

	for {
		select { //selects statement which allows to recieve from whichever channel is ready
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
