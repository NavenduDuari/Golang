package main

import "fmt"

func main(){

	n := 7
	if n%2 == 0 {
		fmt.Println(n," is even\n")
	} else {
		fmt.Println(n," is odd")
	}

	if num := 19; num < 0 {
		fmt.Println(num, "is negative\n")
	} else if num < 10 {
		fmt.Println(num , "has 1 digit\n")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}