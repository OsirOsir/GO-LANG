//Structs Equality

package main

import "fmt"

type infor struct {
	firstName string
	lastName  string
}

func main() {
	P1 := infor{
		firstName: "Philip",
		lastName:  "Osir",
	}

	P2 := infor{
		firstName: "Philip",
		lastName:  "Osir",
	}

	if P1 == P2 {
		fmt.Println("The names are the same ")
	} else {
		fmt.Println("Are not the same")
	}

	P3 := infor{
		firstName: "Philip",
		lastName:  "Osir",
	}

	P4 := infor{
		firstName: "Philip",
	}

	if P3 == P4 {
		fmt.Println("The names are the same ")
	} else {
		fmt.Println("Are not the same")
	}
}
