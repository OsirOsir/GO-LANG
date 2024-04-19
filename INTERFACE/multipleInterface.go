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
	tatalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s salary is $%d\n", e.firstName, e.lastName, e.basicPay)
}

func (e Employee) CalculateLeavesLeft() int {
	return e.tatalLeaves - e.leavesTaken
}

func Leavesremaining(e []Employee) {
	lft := 0
	for _, v := range e {
		lft = v.CalculateLeavesLeft()
	}
	fmt.Printf("The number of leave days remaining is %d\n", lft)
}
func main() {
	emp1 := Employee{
		firstName:   "Philip",
		lastName:    "Osir",
		basicPay:    39000,
		pf:          75,
		tatalLeaves: 39,
		leavesTaken: 24,
	}
	var ds SalaryCalculator
	ds = emp1
	ds.DisplaySalary()

	leaves := []Employee{emp1}
	Leavesremaining(leaves)

}
