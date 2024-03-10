package main

import "fmt"

func uniqueNames(a, b []string) []string {
	// Create a map to store unique names
	unique := make(map[string]bool)

	// Iterate over slice 'a' and add names to the map
	for _, name := range a {
		unique[name] = true
	}

	// Iterate over slice 'b' and add names to the map
	for _, name := range b {
		unique[name] = true
	}

	// Create a new slice to store the unique names
	var result []string

	// Append unique names from the map to the result slice
	for name := range unique {
		result = append(result, name)
	}

	return result
}

func main() {
	// should print Ava, Emma, Olivia, Sophia

	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))

}
