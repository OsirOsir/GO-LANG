package main

import "fmt"

func main() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length  of slice is %d and capacity is %d\n", len(fruitslice), cap(fruitarray))

}
