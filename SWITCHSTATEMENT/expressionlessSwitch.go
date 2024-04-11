package main

import "fmt"

func main() {
	num := 75

	switch { //expression is omitted
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50\n", num)

	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 51 and less than 100\n", num)
	case num >= 101:
		fmt.Printf("%d is greater than 100\n", num)
	}
}
