package main

import "fmt"

func main() {
	x := 0

	for x < 1000 {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			x++
			break
		}
	}
}
