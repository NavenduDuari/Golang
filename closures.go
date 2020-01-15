package main

import (
	"fmt"
)

func outerFunc() func() string {
	return func() string {
		return "This is coming from inner function"
	}
}

func intSeqFunc() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	innerFunc := outerFunc()
	fmt.Println(innerFunc())

	nxtInt := intSeqFunc()
	fmt.Println(nxtInt())
	fmt.Println(nxtInt())
	fmt.Println(nxtInt())

	newNxtInt := intSeqFunc()
	fmt.Println(newNxtInt())
	fmt.Println(newNxtInt())
}
