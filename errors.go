package main

import (
	"errors"
	"fmt"
)

func f1(n int) (int, error) {
	if n == 0 {
		return -1, errors.New("Can't work with 0")
	} else {
		return n + 200, nil
	}
}

type customErrorType struct {
	val int
	msg string
}

// func (e *customErrorType) Error() string {
// 	return fmt.Sprintf("%d-%s", e.val, e.msg)
// }

func f2(n int) (int, *customErrorType) {
	if n == 0 {
		e := &customErrorType{val: n, msg: "Can't work with it"}
		return -1, e
	} else {
		return n + 200, nil
	}
}

func main() {
	for _, n := range []int{0, 1} {
		r, e := f1(n)
		if e != nil {
			fmt.Println("f1 faild:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, n := range []int{0, 1} {
		r, e := f2(n)
		if e != nil {
			fmt.Println("f2 failed:", e.msg)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
}
