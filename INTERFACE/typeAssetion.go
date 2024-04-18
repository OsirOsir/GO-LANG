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
	distance float64
	fuel     float64
}

func (b BMW) Mileage() float64 {
	return b.distance / b.fuel
}

func (a Audi) Mileage() float64 {
	return a.distance / a.fuel
}

func totalMileage(m []MotorVehicle) {
	au := m.(BMW)

	fmt.Printf(au.averagespeed())
}

func main() {
	b1 := BMW{
		distance:     167.9,
		fuel:         45,
		averagespeed: "78",
	}

	// person := []MotorVehicle{b1, a1}

	totalMileage(b1)
}
