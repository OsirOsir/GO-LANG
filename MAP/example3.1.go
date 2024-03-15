package main

import "fmt"

func main() {
	employeSalaries := map[string]int{
		"Joshua":   3000,
		"Alice":    2000,
		"Fredrick": 4000,
	}

	name := "Joshua"
	salary := employeSalaries[name]
	fmt.Println(name, "salary is ", salary)
}
