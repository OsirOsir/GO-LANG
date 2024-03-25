// Declaring struct

package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {
	// creating a struct specifying field names

	emp1 := Employee{
		firstName: "Sam",
		lastName:  "Brown",
		age:       25,
		salary:    5000,
	}

	//creating struct without specifyng field names

	emp2 := Employee{"Andrew", "Johns", 25, 3400}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
}
