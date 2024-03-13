package main

import "fmt"

func main() {
	namesGC := []string{"Aggy", "Zippy", "Alice", "Cathy", "NaomiJ", "NaomiB"}
	fmt.Println("Names of ladiiesCousines", namesGC, len(namesGC), cap(namesGC))

	modNamesGC := namesGC[:]
	modNamesGC = append(modNamesGC, "NaomiW")
	fmt.Println(modNamesGC, len(modNamesGC), cap(modNamesGC))

	specName := modNamesGC[1:4]
	fmt.Println("Name per house are", specName, cap(specName))

	// reslicing
	specName = specName[:cap(specName)]
	fmt.Println(specName, cap(specName))
}
