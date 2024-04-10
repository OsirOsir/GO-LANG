package main

import "fmt"

func main() {
	letter := "j"

	fmt.Printf("Lettter %s is a", letter)

	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vowels \t")
	default:
		fmt.Println("\tNot a vowel")
	}
}
