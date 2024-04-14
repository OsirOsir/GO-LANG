package main

import "fmt"

func main() {
	// letter := 'e'
	// fmt.Printf("%c is ", letter)
	// switch letter {
	// case 'a', 'e', 'i', 'o', 'u':
	// 	fmt.Println("a vowel")
	// default:
	// 	fmt.Println("not a vowel ")
	// }

	num := 102

	switch {
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50\n", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is great than 51 and less than 100\n", num)
	case num > 101:
		fmt.Printf("%d is greater than 100\n", num)
	}
}
