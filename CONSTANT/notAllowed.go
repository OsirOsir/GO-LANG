package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) {
	a := math.Sqrt(n)
	fmt.Printf("Squre root of a is %.1f\n", a)
}

func main() {
	v := 4.0

	sqrt(v)
}
