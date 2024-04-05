package main

import "fmt"

type employee struct {
	name1  string
	name2  string
	age    int
	weight float64
	height float64
}

func main() {
	emp1 := employee{
		name1:  "Alice",
		name2:  "Adhiambo",
		age:    26,
		weight: 67.2,
		height: 175.5,
	}
	fmt.Println("Name1", emp1.name1)
	fmt.Println("Name2", emp1.name2)
	fmt.Println("Age", emp1.age)
	emp1.name1 = "Iddah"
	fmt.Println("Updated details of employee1", emp1)
}
