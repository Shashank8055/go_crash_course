package main

import "fmt"

func main() {
	// Define map
	//emails := make(map[string]string)

	// Assign kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Shashank"] = "shashank@gmail.com"

	// Declare map and Key values
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com", "Shashank": "shashank@gmail.com"}

	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	// Length of map
	fmt.Println(len(emails))
	// To access a particular element using key
	fmt.Println(emails["Bob"])
	// To delete from map
	delete(emails, "Bob")

	fmt.Println(emails)
}
