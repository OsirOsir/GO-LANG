package main

import "fmt"

type occupation struct {
	office    string
	outoffice string
}

type artist struct {
	name   string
	gender string
	occupation
}

func main() {
	P1 := artist{
		name:   "Kyle",
		gender: "Male",
		occupation: occupation{
			office:    "Finance",    // this are promoted fields
			outoffice: "Footballer", // this are promoted fields
		},
	}

	fmt.Println("Type of office Occupation = ", P1.occupation.office)
}
