package main

import "fmt"

type Vehicle interface {
	mileage() float64
}

type SuzukiAlto struct {
	distance     float64
	averagespeed string
}

type BMW struct {
	distance     float64
	averagespeed string
}

type Demio struct {
	distance float64
}

func (s SuzukiAlto) mileage() float64 {
	return s.distance
}

func (b BMW) mileage() float64 {
	return b.distance
}

func (d Demio) mileage() float64 {
	return d.distance
}

func averageDistance(v []Vehicle) {
	TD := 0.0

	for _, r := range v {
		TD = TD + r.mileage()
	}
	AD := TD / float64(len(v))

	fmt.Printf("The Avarage Distance Coverd by the Vehicles is %.1fkm\n", AD)
}

func main() {
	v1 := SuzukiAlto{
		distance:     235,
		averagespeed: "89",
	}
	v2 := BMW{
		distance:     148,
		averagespeed: "70",
	}
	v3 := Demio{
		distance: 179,
	}

	sliceVehicle := []Vehicle{v1, v2, v3}

	averageDistance(sliceVehicle)
}
