package main

import "fmt"

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeave() int
}

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

type Employee struct {
	firstName  string
	lastName   string
	basicPay   int
	pf         int
	totalLeave int
	leaveTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s salary is $%d\n", e.firstName, e.lastName, e.basicPay)
}

func (e Employee) CalculateLeave() int {
	return e.totalLeave - e.leaveTaken
}

func main() {
	emp := Employee{
		firstName:  "Philip",
		lastName:   "Osir",
		basicPay:   30000,
		pf:         5,
		totalLeave: 30,
		leaveTaken: 12,
	}

	var ds EmployeeOperations = emp
	ds.DisplaySalary()

	var lft EmployeeOperations = emp
	fmt.Printf("Leave days remaining is %d\n", lft.CalculateLeave())
}
