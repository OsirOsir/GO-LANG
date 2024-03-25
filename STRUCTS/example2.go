package main

import "fmt"

type persDetails struct {
	firstName string
	lastName  string
	age       int
	height    int
	weight    float64
}

func main() {
	person1 := persDetails{
		firstName: "Iddah",
		lastName:  "Awuor",
		age:       27,
		height:    165,
		weight:    45.5,
	}
	fmt.Println("First Name:", person1.firstName)
	person1.lastName = "Otieno"

	fmt.Printf("First Name: %s  Last Name: %s\n Age: %d\n Height: %d\n Weight: %.1f\n", person1.firstName, person1.lastName, person1.age, person1.height, person1.weight)
}
