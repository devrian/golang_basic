package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

// hasBirthday method (pointer, receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	// init person using struct
	person1 := Person{
		firstName: "Devrian",
		lastName:  "Ansyah",
		city:      "Jkt",
		gender:    "M",
		age:       23,
	}

	// Alternative
	person2 := Person{"Her", "Unknows", "Unk", "F", 20}

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)
	person1.hasBirthday()
	person1.getMarried("Williams")

	person2.getMarried("Ansyah")
	fmt.Println(person2.greet())
}
