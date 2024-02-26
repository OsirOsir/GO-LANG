// 5. **Array Sum:**
//    - Declare an array named `expenses` containing 5 float64 values representing monthly expenses.
//    - Calculate and print the total sum of expenses.

package main

import "fmt"

func main() {
	expenses := [5]float64{102.9, 243.4, 122.6, 566.4, 321.7}
	var sum float64
	for i, v := range expenses {
		fmt.Printf("\n%d exepense value is %.1f\n", i, v)
		sum += v
	}
	fmt.Printf("\nThe total sum is %.3f\n", sum)
}
