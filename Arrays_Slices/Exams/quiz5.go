// 5. **Array Sum:**
//    - Declare an array named `expenses` containing 5 float64 values representing monthly expenses.
//    - Calculate and print the total sum of expenses.

package main

import "fmt"

func main() {
	expenses := [5]float64{224.5, 345.9, 233.7, 102.8, 156.9}

	var sum float64

	for i, v := range expenses {
		fmt.Printf("%d nth element expence is: %.1f\n", i, v)
		sum += v
	}
	fmt.Printf("\nTotal sum of Expenses is: %.1f\n", sum)
}
