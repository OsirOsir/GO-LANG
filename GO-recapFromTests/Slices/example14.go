package main

import "fmt"

func main() {
	Toyota := [][]string{
		{"Landcruiser", "Axio", "PremioV"},
		{"Raum", "Hilux", "Belta"},
		{"NZE", "Caldina", "Filder"},
	}

	Toyota = append(Toyota, []string{"Ractis", "Harriah", "Wish"})
	fmt.Println(Toyota)
	fmt.Println()

	for _, r := range Toyota {
		for _, c := range r {
			fmt.Printf("%s ", c)
		}
		fmt.Printf("\n")
	}
}
