package main

import(
	"fmt"
)

func main(){

	// nums := []int{10,20,30}
	// sum := 0
	// for index, value := range nums {  //range on arrays and slices 
	// 	fmt.Println(index,"=",value)  //provides both the index and value for each entry.
	// }

	m := map [string] string {"a":"apple", "b":"ball"}
	for key, value := range m {
		fmt.Println(key,"for", value)
	}

	
}