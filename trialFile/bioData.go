package main

import "fmt"

type bioData struct {
	firstName  string
	secondName string
	age        int
	gender     string
	salary     int
}

func newData() bioData {
	b := bioData{
		firstName:  "",
		secondName: "",
		age:        0,
		gender:     "",
		salary:     0,
	}

	return b
}

func (b bioData) format() string {
	fs := "Employees Summary\n"

	fs += fmt.Sprintf("Names: %s %s\nAge: %d\nGender: %s\nSalary: %d", b.firstName, b.secondName, b.age, b.gender, b.salary)

	return fs
}

func (b *bioData) updatebioFir(firstName string) {
	b.firstName = firstName
}
func (b *bioData) updatebioseco(secondName string) {
	b.secondName = secondName
}

func (b *bioData) updatebioAge(age int) {
	b.age = age
}

func (b *bioData) updatebioGender(gender string) {
	b.gender = gender
}
func (b *bioData) updatebioSalary(salary int) {
	b.salary = salary
}
