package main

import "fmt"

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalcultaor interface {
	CalculateLeaveLeft() int
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
	fmt.Printf("%s %s Salary is $ %d\n", e.firstName, e.lastName, e.basicPay)
}

func (e Employee) CalculateLeaveLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func leavedayslft(s []LeaveCalcultaor) {

	for _, v := range s {
		v.CalculateLeaveLeft()
	}
	fmt.Printf("Number of leave left is %v\n", v.CalculateLeaveLeft)
}

func main() {
	e1 := Employee{
		firstName:   "Philip",
		lastName:    "Otieno",
		basicPay:    36000,
		pf:          123,
		totalLeaves: 35,
		leavesTaken: 25,
	}
	var sl SalaryCalculator
	sl = e1
	sl.DisplaySalary()

	leavedayslft([]LeaveCalcultaor{e1})
}
