package main

import "fmt"

type employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {
	emp := employee{
		firstName: "Brian",
		lastName:  "Onger",
	}

	fmt.Printf("First Name: %s Last Name: %s\n Age: %d\n Salary: %v", emp.firstName, emp.lastName, emp.age, emp.salary)

}
