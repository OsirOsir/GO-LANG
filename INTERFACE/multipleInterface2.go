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
	fmt.Printf("%s salary is %d\n", e.firstName, e.lastName, e.basicPay)
}

func (e Employee) CalculateLeavesLeft() int {
	return e.tatalLeaves - e.leavesTaken
}

func main() {
	e := Employee{
		firstName:   "Alice",
		lastName:    "Otieno",
		basicPay:    30000,
		pf:          5,
		tatalLeaves: 25,
		leavesTaken: 18,
	}

	ds := e
	ds.DisplaySalary()

	l := e
	fmt.Printf("Remaining leave days are %d", l.CalculateLeavesLeft)
}
