package main

import (
	"fmt"
	"reflect"
) 

func main(){

	//type is auto-assigned if we do not specify it explicitly
	var a = 10
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a)) //int

	var b = 10.5
	fmt.Println(reflect.TypeOf(b))  //float64
	
	var c = "Navendu"
	fmt.Println(reflect.TypeOf(c))  //string

	var d = true
	fmt.Println(reflect.TypeOf(d))  //bool

	//we can explicitly specify datatype
	var e int = 2
	fmt.Println(reflect.TypeOf(e))  //int

	var f string = "Hello world"
	fmt.Println(reflect.TypeOf(f))  //string

	// we can initialize more than 1 variable
	var n1 , n2 int = 4,5
	fmt.Println(n1,n2)

	//If we do not initialize a variable then it is assigned with zero
	var g int
	fmt.Println(g)  //0 

	var h string
	fmt.Println(h)  //" "

	var i float64
	fmt.Println(i)  //0

	//Shorthand for declearing and initializing a variable
	
	j := "apple"  //instead of var j string = apple
	fmt.Println(j)




}