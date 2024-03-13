package main

import "fmt"

func main() {
	slice := make([]int, 5, 5)

	fmt.Println(slice, len(slice), cap(slice))
}
