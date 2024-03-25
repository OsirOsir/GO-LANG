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
	}

	fmt.Println(emp)

}
