package main

import "fmt"

func mainManu(y map[string]float64) {
	y["coffee"] = 2.99
	y["spaniad coktail"] = 3.66
}

func main() {
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}
	mainManu(menu)
	fmt.Println(menu)
}
