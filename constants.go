package main

import(
	"fmt"
	"reflect"
	"math"
)

const s string = "Constant"

func main(){
	fmt.Println(s)

	const n = 50000000
	fmt.Println(reflect.TypeOf(n))

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}