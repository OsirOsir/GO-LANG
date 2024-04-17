package main

import "fmt"

type MotorVehicle interface {
	Mileage() float64
}

type BMW struct {
	distance     float64
	fuel         float64
	averagespeed string
}

type Audi struct {
	distance     float64
	fuel         float64
	averagespeed string
}

type Subaru struct {
	distance float64
	fuel     float64
}

func (b BMW) Mileage() float64 {
	return b.distance / b.fuel
}

func (a Audi) Mileage() float64 {
	return a.distance / a.fuel
}

func (s Subaru) Mileage() float64 {
	return s.distance / s.fuel
}

func totalMileage(m []MotorVehicle) {
	var tm float64

	for _, v := range m {
		tm += v.Mileage()
	}

	fmt.Printf("The total mileage per month is %.5f km/l\n", tm)
}

func main() {
	bi := BMW{
		distance:     178.1,
		fuel:         36,
		averagespeed: "87",
	}
	a1 := Audi{
		distance:     134.9,
		fuel:         45,
		averagespeed: "90",
	}

	s1 := Subaru{
		distance: 178.3,
		fuel:     56,
	}

	person1 := []MotorVehicle{bi, a1, s1}

	totalMileage(person1)
}
