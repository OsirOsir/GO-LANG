package main

import "fmt"

func main() {
	animals := [3][3]string{
		{"Tiger", "Leoperd", "Giraffe"},
		{"Lion", "Kangaroo", "Zebra"},
		{"Buffalo", "Dog", "Gazelle"},
	}

	domAnimals := animals
	domAnimals[1] = [3]string{"Chicken", "Cow", "Goat"}

	// fmt.Println(animals, len(animals))

	for r := 0; r <= 2; r++ {
		for c := 0; c <= 2; c++ {
			fmt.Println(domAnimals[r][c])
		}
	}

}
