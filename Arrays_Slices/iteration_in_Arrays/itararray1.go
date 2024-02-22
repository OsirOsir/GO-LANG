package main

import "fmt"

func main() {
	num := [...]float64{34.5, 45.7, 89.9, 56.4, 34.1, 90.5}

	for i := 0; i < len(num); i++ {
		fmt.Printf("%d th element of num is: %0.1f\n", i, num[i])
	}
}
