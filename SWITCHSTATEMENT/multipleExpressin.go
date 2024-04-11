package main

import "fmt"

func main() {
	letter := 'k'

	fmt.Printf("Letter %c ", letter)

	switch letter {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("is  a vowel")
	default:
		fmt.Println("not a vowel")
	}

}
