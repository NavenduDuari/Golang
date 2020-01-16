package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan string, results chan<- string) {
	for {
		j, more := <-jobs
		if more {
			fmt.Printf("Worker %d Received: %s \n", id, j)
			time.Sleep(time.Second * 2)
			results <- fmt.Sprintf("Worker %d Result: %s", id, j)
		} else {
			fmt.Printf("Worker %d : All jobs received \n", id)
			break
		}
	}
}

func main() {

	jobs := make(chan string, 10)
	results := make(chan string, 10)

	go worker(1, jobs, results)
	go worker(2, jobs, results)
	go worker(3, jobs, results)
	go worker(4, jobs, results)

	for i := 1; i <= 10; i++ {
		jobs <- fmt.Sprintf("job:%d", i)
		fmt.Println("jobs being sent..")
	}
	close(jobs)

	for i := 1; i <= 10; i++ {
		fmt.Println(<-results)
	}

}
