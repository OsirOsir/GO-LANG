package main

import "fmt"

type name struct {
	firstName string
	lastName  string
}

func main() {
	data1 := name{
		firstName: "Philip",
		lastName:  "Osir",
	}

	data2 := name{
		firstName: "Opon",
		lastName:  "Joshua",
	}

	if data1 == data2 {
		fmt.Println("Names are similer")
	} else {
		fmt.Println("Names are not same")
	}
}
