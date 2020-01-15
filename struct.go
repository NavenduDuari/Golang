package main

import (
	"fmt"
	// "reflect"
)

type person struct {
	name string
	age int
}


func main(){

	p1 := person{"Navendu", 22}
	fmt.Println(p1.name)

	p2 := person{name: "Puja", age: 20}
	fmt.Println(p2.age)

	p1.age = 25
	fmt.Println(p1.age)

	p3 := p1
	fmt.Println(p3.name)

	p3.name = "Bob"
	fmt.Println("p1.name: ",p1.name,"\np3.name:",p3.name)

	p1add := &p1
	p1add.name = "Duari"
	fmt.Println("p1.name: ",p1.name,"p1add.name: ",p1add.name)
}