package main

import (
	"fmt"
)

func main(){

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:",m)

	m["k3"] = 99
	fmt.Println(m)

	delete(m, "k2")
	fmt.Println(m)

	v1 := m["k1"]    //get a value for key with name[key]
	fmt.Println("v1:",v1)

	fmt.Println(len(m))  //# of key value pair

	val,prs := m["k3"]   //return value , presence
	fmt.Println(val,prs)

	n := map[string]int{
		"Navendu": 22,
		"Puja": 22,
	}
	fmt.Println(n)
	
}