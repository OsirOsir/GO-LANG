package main

import "fmt"

func main() {

	// Sort hand Declaration
	month := [...]string{"January", "February", "March", "April", "May"} // we can ignore the length of the array
	// in the declaration and replace it with  ... and let the compiler find the length fo us

	// for _, mon := range month {
	// 	fmt.Println(mon)
	// }
	fmt.Println(month, len(month))

	// fmt.Println(month)
}
