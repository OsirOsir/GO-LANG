package main

import "fmt"

type MotorVehicle interface {
	Mileage() float64
}

type BMW struct {
	distance     float64
	fuel         float64
	averagespeed int
}

func (b BMW) Mileage() float64 {
	return b.distance / b.fuel
}

func AvarageSpeed(m []MotorVehicle) {
	for _, vehicle := range m {
		if BMW, ok := vehicle.(BMW); ok {
			fmt.Println(BMW.averagespeed)
		}
	}
}

func main() {
	bV := BMW{
		distance:     167.9,
		fuel:         45,
		averagespeed: 98,
	}

	AvarageSpeed([]MotorVehicle{bV})
}
