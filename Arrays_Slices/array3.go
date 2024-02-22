package main

import "fmt"

func main() {
	// How we can manupulate  arrays  ,  but it doesnt change the first array
	// It changes the value of that particular index
	names := [...]string{"Philip", "Otieno", "Ogutu", "Janyakach", "Kadianga,"}

	upNames := names

	upNames[3] = "Koguta"

	for _, cuNames := range upNames {
		fmt.Println(cuNames)
	}

	fmt.Println("The Original names are: ", names)
	// fmt.Println("The new Nmaes are: ", upNames)
}
