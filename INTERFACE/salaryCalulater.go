package main

import "fmt"

type salaryCalcultar interface {
	companySalary() int
}

type PamanentEmployees struct {
	emp1 int
	emp2 int
	emp3 int
	emp4 int
}

type ContractedEmployees struct {
	emp1 int
	emp2 int
	emp3 int
}

func (p PamanentEmployees) companySalary() int {
	return p.emp1 + p.emp2 + p.emp3 + p.emp4
}

func (c ContractedEmployees) companySalary() int {
	return c.emp1 + c.emp2 + c.emp3
}

func totalExpence(s []salaryCalcultar) {
	te := 0

	for _, v := range s {
		te += v.companySalary()
	}
	fmt.Printf("Total company expences is : USD %d\n", te)
}

func main() {
	pp := PamanentEmployees{
		emp1: 36000,
		emp2: 45988,
		emp3: 24556,
		emp4: 37888,
	}
	cc := ContractedEmployees{
		emp1: 98666,
		emp2: 45777,
		emp3: 78999,
	}

	totalExpence([]salaryCalcultar{pp, cc})
}
