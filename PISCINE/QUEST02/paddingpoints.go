package main

import "fmt"

func main() {
	// Printing integers with padding
	num := 123

	fmt.Printf("Padded with spaces:%5d\n", num)  //Prints: Padded with spaces:     123
	fmt.Printf("Zero-Padded:%05d\n", num)        // Prints: Zero-padded: 00123
	fmt.Printf("Padded with spaces: %3d\n", num) // Prints: Padded with spaces: 123
	fmt.Printf("Zero-Padded: %03d\n", num)       //Prints: Zero-padded: 123

	// Negative numbers with padding
	negativeNum := -45
	fmt.Printf("Padded with spaces: %6d\n", negativeNum) // Prints: "Padded with spaces:    -45"
	fmt.Printf("Zero-padded: %06d\n", negativeNum)       // Prints: "Zero-padded: -00045"

	// Printing multiple integers with padding
	fmt.Printf("%5d %05d\n", 123, 456)  // Prints: "  123 00456"
	fmt.Printf("%-5d %05d\n", 123, 456) // Prints: "123   456"
}
