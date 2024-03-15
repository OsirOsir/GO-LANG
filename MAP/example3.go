package main

import "fmt"

func main() {
	resultsMarks := map[string]float64{
		"Maths":     45.7,
		"Kiswahili": 65.8,
		"English":   89.3,
	}
	results := "Maths"

	fmt.Println(results, "is", resultsMarks["Maths"])
}
