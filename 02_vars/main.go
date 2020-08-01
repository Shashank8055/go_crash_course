package main

import "fmt"

func main() {
	// Main Types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint int8 uint16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	//var name string = "Shashank"

	var age int = 37
	const isCool = true

	// Shorthand
	//name := "Shashank"
	//email := "shashankpatel8055@yahoo.com"
	height := 5.9

	name, email := "Shashank Patel", "shashankpatel8055@yahoo.com"

	//fmt.Println(age)
	//fmt.Println(name)
	fmt.Println(name, email, age, height, isCool)
	//fmt.Printf("&T", name)
	//fmt.Printf("%T", age)

}
