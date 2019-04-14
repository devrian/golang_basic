package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign key value
	// emails["Dev"] = "dev@gmail.com"
	// emails["Ria"] = "ria@gmail.com"
	// emails["Syah"] = "syah@gmail.com"

	// Delare map and add key value
	emails := map[string]string{"Dev": "dev@gmail.com", "Ria": "ria@gmail.com"}
	emails["Syah"] = "syah@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Dev"])

	// Delete from map
	delete(emails, "Dev")
	fmt.Println(emails)
}
