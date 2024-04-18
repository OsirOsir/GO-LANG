package main

import "fmt"

type SalaryCalcultar interface {
	calculateSalary() float64
}

type Pamanent struct {
	empID     string
	bSalary   float64
	allwance  float64
	dedaction float64
}

type Contracted struct {
	empID   string
	bSalary float64
}

type freelancer struct {
	empID     string
	ratePerHr float64
	totalHrs  float64
}

func (p Pamanent) calculateSalary() float64 {
	return (p.bSalary + p.allwance) - p.dedaction
}

func (c Contracted) calculateSalary() float64 {
	return c.bSalary
}
func (f freelancer) calculateSalary() float64 {
	return f.ratePerHr * f.totalHrs
}
func totalExpence(s []SalaryCalcultar) {
	te := 0.0
	for _, value := range s {
		te += value.calculateSalary()
	}

	fmt.Printf("Total Employee Expences is USD %.2f\n", te)
}

func main() {
	emp1 := Pamanent{
		empID:     "Philip Osir",
		bSalary:   36999.33,
		allwance:  8000.9,
		dedaction: 400.88,
	}

	emp2 := Pamanent{
		empID:     "Kyle Hope",
		bSalary:   26800.3,
		allwance:  5000.9,
		dedaction: 400.88,
	}

	emp3 := Contracted{
		empID:   "Iddah Awuor",
		bSalary: 56777.2,
	}

	fl := freelancer{
		empID:     "Alice Adhiambo",
		ratePerHr: 3600.9,
		totalHrs:  56,
	}

	employee := []SalaryCalcultar{emp1, emp2, emp3, fl}

	totalExpence(employee)
}
