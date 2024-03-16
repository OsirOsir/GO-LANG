package main

import "fmt"

type Employee struct {
	salary  int
	country string
}

func main() {
	employees := map[string]Employee{
		"Getrude": {salary: 30000, country: "USA"},
		"Alice":   {salary: 40000, country: "Kenya"},
		"Wendy":   {salary: 35000, country: "Uganda"},
	}

	for k, v := range employees {
		fmt.Printf("Employee: %s, Salary: %d, Country: %s\n", k, v.salary, v.country)
	}
}
