package main

import "fmt"

// Closing a channel indicates that no more values
// will be sent on it. This can be useful to communicate
// completion to the channel’s receivers.

func boss(jobs chan<- string) {
	for i := 1; i <= 5; i++ {
		jobs <- fmt.Sprintf("Job %d", i)
		fmt.Println("Job sent")
	}
	// close(jobs)
	fmt.Println("All jobs sent")
}

func employee(jobs <-chan string, done chan<- bool) {
	for {
		// j, more := <-jobs
		// if more {
		// 	fmt.Println("job received:", j)
		// } else {
		// 	fmt.Println("All jobs received")
		// 	done <- true
		// 	break
		// }
		j := <-jobs // why this does not cause Deadlock
		//while channel is not closed
		fmt.Println(j)
	}
}

func main() {

	jobs := make(chan string)
	done := make(chan bool)

	go boss(jobs)
	// for {         //this results deadlock if channel is not closed
	// 	j := <-jobs
	// 	fmt.Println(j)
	// }
	go employee(jobs, done)

	fmt.Scanln()
}
