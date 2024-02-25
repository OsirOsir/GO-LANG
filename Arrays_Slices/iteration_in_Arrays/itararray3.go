package main

import "fmt"

func main() {
	price := [...]float64{68.9, 43.4, 56.8, 45.1}
	var sum float64
	for i, v := range price {
		fmt.Printf("%d th value of the price is %.1f\n", i, v)
		sum = v + sum
	}

	fmt.Printf("\nsum of all Prices is: %.1f\n", sum)
}
