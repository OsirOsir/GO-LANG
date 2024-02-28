// **Question: Array Rotation**

// - Declare an array named `numbers` containing 7 integers.
// - Write a function to rotate the elements of the `numbers` array to the left by a given number of positions and print the rotated array.

package main

import "fmt"

func rotateLeft(numbers []int, positions int) {
	length := len(numbers)
	positions %= length // Ensure positions is within the length of the array

	reverse(numbers[:positions])
	reverse(numbers[positions:])
	reverse(numbers)
}

func reverse(numbers []int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	rotateLeft(numbers, 2) // Rotate left by 2 positions

	fmt.Println(numbers) // Output: [3 4 5 6 7 1 2]
}

// package main

// import "fmt"

// func rotateNumbers(numbers []int) {
// 	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
// 		numbers[i], numbers[j] = numbers[j], numbers[i]
// 	}
// }

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5, 6, 8}

// 	rotateNumbers(numbers)

// 	fmt.Println(numbers)
// }
