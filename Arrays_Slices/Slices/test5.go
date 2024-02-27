package main

import "fmt"

func main() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}

	fruitslice := fruitarray[1:3]
	fmt.Println(fruitarray, len(fruitarray))
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))

}
