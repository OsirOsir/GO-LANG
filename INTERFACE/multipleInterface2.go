package main

import "fmt"

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s salary is %d\n", e.firstName, e.lastName, e.basicPay)
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := Employee{
		firstName:   "Alice",
		lastName:    "Otieno",
		basicPay:    30000,
		pf:          5,
		totalLeaves: 25,
		leavesTaken: 18,
	}

	var ds SalaryCalculator = e
	ds.DisplaySalary()

	var l LeaveCalculator = e
	fmt.Printf("Remaining leave days are %d\n", l.CalculateLeavesLeft())
}
