package main

import "fmt"

type MotorVehicle interface {
	mileage() int
}

type Subaru struct {
	distance int
	fuel     int
	avespeed string
}

func (s Subaru) mileage() int {
	return s.distance / s.fuel
}

func averageSpeed(m []MotorVehicle) {
	for _, v := range m {
		fmt.Printf("Mileage is %d\n", v.mileage())
		if sub, ok := v.(Subaru); ok {
			fmt.Printf("The average speed of Subaru is %s\n", sub.avespeed)
		}
	}
}

func main() {
	s1 := Subaru{
		distance: 34,
		fuel:     5,
		avespeed: "67",
	}
	avsp := []MotorVehicle{s1}
	averageSpeed(avsp)
}
