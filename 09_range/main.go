package main

import "fmt"

func main() {
	ids := []int{11, 22, 34, 43, 22, 1}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - Id: %d\n", i, id)
	}

	// Not using index
	for id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Dev": "dev@gmail.com", "Ria": "ria@gmail.com"}

	for key, val := range emails {
		fmt.Printf("%s: %s\n", key, val)
	}

	for key := range emails {
		fmt.Println("Name: " + key)
	}
}
