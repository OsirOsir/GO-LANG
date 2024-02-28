//Passing a slice to a function

package main

import "fmt"

func subtractOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func main() {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)

	subtractOne(nos)
	fmt.Println("slice after function call", nos)
}
