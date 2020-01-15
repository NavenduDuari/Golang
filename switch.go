package main

import (
	"fmt"
	"time"
	"reflect"
)

func main(){

	i := 2
	fmt.Println("write",i," as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday :
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//switch without an expression is an alternate way 
	//to express if/else logic. Here we also show how 
	//the case expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}


	fmt.Println(reflect.TypeOf(time.Sunday))

	fmt.Println(4<5 == true)

	fmt.Println(reflect.TypeOf(time.Sunday))

	fmt.Println(time.Now().Weekday())
}