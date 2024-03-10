package main

func incrementBy(x int) (f func(int) int) {
	return func(input int) int {
		return input + x
	}
}
