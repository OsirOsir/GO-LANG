package main

import "fmt"

func find(mainName string, Othernames ...string) {
	fmt.Printf("type of othernames is %T\n", Othernames)

	found := false
	for i, v := range Othernames {
		if v == mainName {
			fmt.Println(mainName, "is found at index", i, "in ", Othernames)
			found = true
		}
	}
	if !found {
		fmt.Println(mainName, "not found in ", Othernames)
	}
	fmt.Printf("\n")
}
