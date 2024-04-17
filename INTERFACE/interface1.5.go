package main

import "fmt"

type Vihicle interface {
	mileage() int
}

type Passo struct {
	distance     int
	fuel         int
	averagespeed string
}

type Demio struct {
	distance     int
	fuel         int
	averagespeed string
}

type SuzukiAlto struct {
	distance int
	fuel     int
}

func (p Passo) mileage() int {
	return p.distance / p.fuel
}

func (d Demio) mileage() int {
	return d.distance / d.fuel
}

func (s SuzukiAlto) mileage() int {
	return s.distance / s.fuel
}
func totalAvarageMileage(v []Vihicle) {
	TAM := 0

	for _, v := range v {
		TAM = TAM + v.mileage()
		// TAM = TAM / 3
	}
	fmt.Printf("The Total avarage mileage of the Vehices is %dkm/l\n", TAM)
}

func main() {
	vP := Passo{
		distance:     167,
		fuel:         56,
		averagespeed: "89",
	}
	vD := Demio{
		distance:     150,
		fuel:         34,
		averagespeed: "67",
	}
	vS := SuzukiAlto{
		distance: 189,
		fuel:     60,
	}

	averageMileage := []Vihicle{vP, vD, vS}

	totalAvarageMileage(averageMileage)
}
