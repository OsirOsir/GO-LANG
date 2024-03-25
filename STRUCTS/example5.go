package main

import "fmt"

type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		string: "naveen",
		int:    50,
	}

	fmt.Println(p1.int)
	fmt.Println(p1.string)

}
