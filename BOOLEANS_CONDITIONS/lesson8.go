package main

import "fmt"

func main() {
	age := 45

	// Boolean exressions
	fmt.Println(age < 45)
	fmt.Println(age > 45)
	fmt.Println(age == 45)
	fmt.Println(age != 45)

	// Conditions
	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is 45")
	}
}
