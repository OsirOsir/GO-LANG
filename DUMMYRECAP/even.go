package main

import "fmt"

func average(a, b float32) float32 {
	return a + b/2
}

func main() {
	var a float32 = 1
	var b float32 = 5
	fmt.Println(average(a, b))
}
