// Modifying Slice

package main

import "fmt"

func main() {
	ages := [...]int{45, 95, 68, 43, 87, 65, 64, 98}

	fmt.Println("Original ages", ages)

	updatedAges := ages[2:6]

	for i := range updatedAges {
		updatedAges[i]++
	}
	fmt.Println("The Updated Ages are: ", ages)
	fmt.Println("Specific new update ages: ", updatedAges)

}
