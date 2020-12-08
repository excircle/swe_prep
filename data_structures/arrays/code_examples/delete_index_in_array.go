package main

import (
	"fmt"
)

//Delete integer value at index 2

func main() {
	//Array indexed from 0 to 3
	a := []int{1, 5, 7, 10}
	
	//Index to target
	index := 2
	
	//VALUE A: Array that has first 2 elements
	fmt.Println(a[:index])
	
	//VALUE B: Last element in array
	fmt.Println(a[len(a)-1])
	
	//Build new array composed of VALUE A + VALUE B
	//Use "..." syntax to tell compiler to join 2 arrays.
	a = append(a[:index], a[index+1:]...)
	
	//ORIGINAL VALUE: [1 5 7 10]
	//NEW VALUE: [1 5 10]
	//Returns new value array with integer 7 missing.
	fmt.Println(a)
}

