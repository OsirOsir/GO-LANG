package main

import "fmt"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		fmt.Println("n should be greater than 0 and less than 10")
		return
	}
	printCombinations("", n, 0)
	fmt.Println() // Print a new line after all combinations for clarity
}

func printCombinations(prefix string, n, start int) {
	if n == 0 {
		fmt.Print(prefix)
		if start < 10-n {
			fmt.Print(", ")
		}
		return
	}
	for i := start; i <= 9; i++ {
		newPrefix := prefix + fmt.Sprintf("%d", i)
		printCombinations(newPrefix, n-1, i+1)
	}
}

func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}
