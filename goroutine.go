package main

import (
	"fmt"
	"time"
)

func f1() {
	for i := 0; i <= 5; i++ {
		fmt.Println("f1:", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func f2() {
	for i := 0; i <= 5; i++ {
		fmt.Println("f2:", i)
		time.Sleep(time.Millisecond * 500)
	}
}
func main() {

	go f1()

	f2()

	fmt.Scanln()
}
