package main

import "fmt"

func sum(n []float64) float64 {
	var total float64

	for _, v := range n {
		total += v
	}
	return total
}

func main() {
	fmt.Println(sum([]float64{5, 6.3, 3, 8, 9, 9, 1}))
}
