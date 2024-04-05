package main

import "fmt"

type person struct {
	string
	int
}

func main() {
	p1 := person{
		string: "Philip",
		int:    34,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
