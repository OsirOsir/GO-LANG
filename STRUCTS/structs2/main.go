package main

import (
	"fmt"
	"structs2/structFile"
)

func main() {
	person1Info := structFile.Details{
		Name:   "Philip Osir",
		Age:    30,
		Gender: "Male",
		Height: 172,
		Weight: 78.9,
		Address: structFile.Address{
			Country:  "Kenya",
			Province: "Nyanza",
		},
	}

	fmt.Printf("Name: %s\nAge: %d\nGender: %s\nAddress-Country: %s\n", person1Info.Name, person1Info.Age, person1Info.Gender, person1Info.Address.Country)
}
