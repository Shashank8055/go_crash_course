package main

import "fmt"

func main() {
	// // Arrays
	// var fruitsArr [2]string

	// // Assign values
	// fruitsArr[0] = "Apple"
	// fruitsArr[1] = "Orange"
	// fmt.Println(fruitsArr)
	// fmt.Println(fruitsArr[1])

	// // Declare and Assign
	// animalsArr := [2]string{"Tiger", "Lion"}
	// fmt.Println(animalsArr)

	// Slice
	fruitSlice := []string{"apple", "orange", "grape", "cherry"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice)) // To calculate the length
	fmt.Println(fruitSlice[1:3]) // Range
}
