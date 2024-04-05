package main

import "fmt"

type employee struct {
	firstName string // firstName is a field
	lastName  string // lastName is also a field
	age       int
	salary    int
}

func main() {
	emp1 := employee{
		firstName: "Joshua",
		lastName:  "Otieno",
		age:       35,
		salary:    30000,
	}
	fmt.Println("employee1", emp1)

	//  we can also creat structs without specifying the fields

	emp2 := employee{"Philip", "Osir", 26, 30000}

	fmt.Println("employee2", emp2)
}
