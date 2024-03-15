package main

import "fmt"

func main() {
	nameAge := map[string]int{
		"Philip": 30,
		"Alice":  26,
	}

	newName := "Dickson"
	value, ok := nameAge[newName]
	if ok == true {
		fmt.Println("The age of ", newName, "is", value)
	} else {
		fmt.Println(newName, "Name not found")
	}
}
