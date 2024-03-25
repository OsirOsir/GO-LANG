// Nested structs

package main

import "fmt"

type Status struct {
	country  string
	province string
}

type bioData struct {
	name   string
	age    int
	gender string
	status Status
}

func main() {
	P1 := bioData{
		name:   "Alice",
		age:    26,
		gender: "Female",
		status: Status{
			country:  "Kenya",
			province: "Nyanza",
		},
	}
	fmt.Println("Country= ", P1.status.country)
	fmt.Println("Name: ", P1.name)
}
