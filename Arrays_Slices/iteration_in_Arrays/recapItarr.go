package main

import "fmt"

func main() {
	fuel := [...]float64{102.4, 154.8, 123.7, 144.5}

	newFuel := fuel
	newFuel[1] = 110.6
	sum := float64(0)
	for i := 0; i < len(newFuel); i++ {
		fmt.Printf("%d nth element of fuel %0.1f\n", i, newFuel[i])
		sum += newFuel[i]
	}
	fmt.Printf("\n The sum of the newFuel list is: %0.1f\n", sum)
}
