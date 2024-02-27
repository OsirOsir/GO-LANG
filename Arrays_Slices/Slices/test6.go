package main

import "fmt"

func main() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]

	fmt.Printf("The length of slice is %d while the capacity is %d\n", len(fruitslice), cap(fruitslice))
	fruitslice = fruitslice[:cap(fruitslice)] //  re-slicing till the capacity

	fmt.Printf("The length of the New re-sliced slice is %d while the final capacity %d\n", len(fruitslice), cap(fruitslice))
}
