package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {

	p := person{
		name: name,
		age:  age,
	}
	return &p
}

func main() {
	p1 := newPerson("Navendu", 22)
	p2 := newPerson("Puja", 21)
	p3 := newPerson("Subhendu", 23)

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3.name, p3.age)
	// fmt.Println(reflect.Type(p3)) //how to see the type of a struct object
}
