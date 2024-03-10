package main

import "fmt"

func main() {
	num := [...]float64{76.4, 71.8, 75, 58.9}

	sum := float64(0)

	for _, v := range num {
		sum += v
	}
	fmt.Printf("%.1f\n", sum)
}
