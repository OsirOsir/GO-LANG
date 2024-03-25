// Nested structs

package main

import "fmt"

type Address struct {
	country  string
	province string
}

type bioData struct {
	name    string
	age     int
	gender  string
	address Address
}

func main() {
	P1 := bioData{
		name:   "Alice",
		age:    26,
		gender: "Female",
		address: Address{
			country:  "Kenya",
			province: "Nyanza",
		},
	}
	fmt.Println("Country= ", P1.address.country)
	fmt.Println("Name: ", P1.name)
}
