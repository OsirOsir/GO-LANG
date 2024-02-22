package main

import "fmt"

func nthage(age [6]int) {
	for i := 0; i < len(age); i++ {
		fmt.Println(i, "th age is", age[i]) // This is the function
	}
}
func main() {
	age := [...]int{5, 34, 9, 10, 14, 4} // This is a test Program
	nthage(age)
}
