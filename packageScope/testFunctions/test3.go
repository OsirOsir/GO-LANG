// Function as a Parameter:

// Write a function called applyFunc that takes a function f as a parameter
// along with an integer x and applies f to x. f should be a function that
// takes an integer as input and returns an integer.

package main

func f(x int) int {
	new := x * 2
	return new
}

func applyFunc(x int, f func(int) int) int {
	return f(x)
}
