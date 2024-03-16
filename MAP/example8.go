// package main

// import "fmt"

// func main() {
// 	countries := make(map[string]int)
// 	countries["Kenya"] = 54
// 	countries["Uganda"] = 67
// 	countries["Tanzania"] = 89
// 	countries["Cameroon"] = 78
// 	countries["Congo"] = 45
// 	countries["Sudan"] = 65

// 	countries["Kenya"] = 80
// 	fmt.Println(countries)

// 	for k, v := range countries {
// 		fmt.Println("The index of", k, "is", v)
// 	}
// 	newCountry := "Burundi"
// 	value, ok := countries[newCountry]
// 	if ok == true {
// 		fmt.Printf("%s is available and its index is %d \n", newCountry, value)
// 	} else {
// 		fmt.Println(newCountry, "not available")
// 	}
// 	// delete(countries, "Kenya")
// 	// fmt.Println(countries)
// }

package main

import "fmt"

type biodata struct {
	age    int
	status string
}

func main() {
	bioStatus := map[string]biodata{
		"Philip":  {34, "Married"},
		"Iddah":   {25, "Single"},
		"Julius":  {28, "Divosed"},
		"Getrude": {35, "Engaged"},
	}
	for k, v := range bioStatus {
		fmt.Println("Nmae is :", k, "age is", v.age, "status", v.status)
	}
}
