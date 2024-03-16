package main

import "fmt"

func main() {
	colors := map[int]string{
		10: "Blue",
		14: "Yellow",
		12: "Red",
	}
	fmt.Println("Colors befor delete", colors)
	delete(colors, 10)
	fmt.Println(colors)
}
