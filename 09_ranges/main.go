package main

import "fmt"

func main() {
	ids := []int{33, 45, 67, 89, 90, 4, 3}

	// Loop through idsd
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Adds ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum =", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Shashank": "shashank@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

}
