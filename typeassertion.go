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
	for _, vehicle := range m {
		fmt.Printf("Mileage: %.2f\n", vehicle.Mileage())
		// Check if the vehicle is a BMW, then print its average speed
		if bmw, ok := vehicle.(BMW); ok {
			fmt.Printf("Average Speed: %s\n", bmw.averagespeed)
		}
	}
}

func main() {
	b1 := BMW{
		distance:     167.9,
		fuel:         45,
		averagespeed: "78",
	}
	vehicles := []MotorVehicle{b1}
	totalMileage(vehicles)
}
