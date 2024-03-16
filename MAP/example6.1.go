package main

import "fmt"

type employee struct {
	salary  int
	country string
}

func main() {
	emp1 := employee{
		salary:  22000,
		country: "USA",
	}
	emp2 := employee{
		salary:  10000,
		country: "Canada",
	}
	emp3 := employee{
		salary:  45000,
		country: "India",
	}
	emp4 := employee{
		salary:  114000,
		country: "Kenya",
	}

	employeeInfo := map[string]employee{
		"steve":  emp1,
		"Jamie":  emp2,
		"Mike":   emp3,
		"Joshua": emp4,
	}

	for name, infor := range employeeInfo {
		fmt.Println("Employee:", name, "Salary:", infor.salary, "Country:", infor.country)
	}
}
