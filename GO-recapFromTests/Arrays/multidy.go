package main

import "fmt"

func printArray(a [3][3]string) {
	for _, r := range a {
		for _, c := range r {
			fmt.Printf("%s ", c)
		}
		fmt.Println()
	}
}

func main() {
	a := [3][3]string{
		{"Grapes", "Ovacado", "Banana"},
		{"Mangoe", "Lemon", "Berries"},
		{"Orange", "Apple", "PawPaw"},
	}

	updated := a

	updated[1][2] = "Guava"
	printArray(updated)

	var b [3][3]string

	b[0][0] = "apple"
	b[0][1] = "sumsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Println()
	printArray(b)
}
