package main

import "fmt"

var ageOne int = 16
var ageTwo = 24

func main() {

	// fmt.Println("Hello, Philip!")

	// Strings

	var nameOne string = "philip"
	var nameTwo = "Alice"
	nameThree := "Iddah"
	var nameFour string

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	var ageThree int
	ageFour := 34

	nameTwo = "Kyle"

	fmt.Println(ageOne, ageTwo, ageThree, ageFour, nameTwo)

	ageThree = 12
	var ageFive int16 = 300

	ageSix := 65555.4098098080

	fmt.Println(ageThree, ageFive, ageSix)

	var numLetter string
	nameFour = "Joshuua"
	nameFour = "J"
	slice := []int{1, 2, 3, 4}
	slice[1] = 10
	fmt.Println(numLetter, nameFour, slice)

	x := 3.14
	y := float64(x)

	fmt.Println(y)

}
