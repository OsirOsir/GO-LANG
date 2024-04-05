package main

import "fmt"

type employee struct {
	name1  string
	name2  string
	age    int
	weight float64
	salary int
}

func main() {
	var emp4 employee

	fmt.Println("Name1", emp4.name1)
	fmt.Println("Name2", emp4.name2)
	fmt.Println("Age", emp4.age)

}
