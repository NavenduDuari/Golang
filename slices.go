package main

import (
	"fmt"
	// "reflect"
)

func main(){

	// var s [] string
	// s = make([]string, 3)
	// fmt.Println(reflect.TypeOf(s))

	// ids := make([]int, 5) //shorthand to initialize
	// fmt.Println(len(ids))

	// fibonacci := []int {1,1,2,3,5,8}
	
	names := []string{
		"Ram",
		"Sam",
		"Jodu",
		// "Modhu", the advantage of this syntax is that we can comment out a particular vlaue
	}

	// fmt.Println(fibonacci[:3])

	// top2Names := names[0:2]
	// top3Names := names[0:3]
	// top2Names[0] = "Rahim"  //this change the value of both  top2Names & top3Names
	// 						//because both are sharig the same array underneath

	// fmt.Println(top2Names)
	// fmt.Println(top3Names)

	//slice of Structs

	book := [] struct{
		id int
		name string
		available bool 
	}{
		{1, "Linux-book", true},
		{2, "Windows-book", false},
		{3, "Mac-book", true},
	}
	fmt.Println(book)

	fmt.Println(len(book))
	fmt.Println(cap(book))

	s := make([] string, 3, 6)
	fmt.Println(len(s))  //length – the number of elements in a slice
	fmt.Println(cap(s))  //capacity – total number of elements in the slice’s underlying array from the first element in the slice.

	//Nil sclice -- A nil slice has a length and capacity of 0 and has no underlying array
	var nilSclice [] string

	if nilSclice == nil {               // nil is equivalent to NULL
		fmt.Println("\n nilSlice is empty")
	}
	
	notNilSclice := make([] string, 0)
	if notNilSclice != nil {
		fmt.Println("\n notNilScilce is not empty")
	}

	//Append to an existing sclice
	names = append(names, "Subhendu", "Shefali")
	fmt.Println(names)

	//Two Dimensional Sclices

	count := 0
	twoD := make([][] int, 4)
	for i := 0; i<4; i++ {
		twoD[i] = make([] int, 2)
		for j := 0; j < 2; j++ {
			twoD[i][j] = count
			count++
		}
	}

	fmt.Println(twoD)
	fmt.Println(twoD[0])  // gives the whole row Here the 1st row

	//LOOP thorough sclice elements
	for index, value := range book {     // range is a go command which will 
		fmt.Println(index, "=", value)  //return the index number and the value of an element.
	}

	for _, name := range names {     //An _ underscore indicates Go language that 
		fmt.Println(name)           //you take the value of a particular variable 
	}                               //(in this case index) and just ignore it.

}