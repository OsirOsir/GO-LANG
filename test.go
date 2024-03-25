package main

import "fmt"

func main() {
	employeSalaries := map[string]int{
		"Philip": 34000,
		"Erick":  50000,
		"Elvis":  90000,
	}

	employeSalaries["Fred"] = 24000

	for k, v := range employeSalaries {
		fmt.Println(k, v)
	}
	newPerson := "Elvis"
	// value, ok := map[key]
	value, ok := employeSalaries[newPerson]

	if ok == true {
		fmt.Printf("Slary of %s is %v\n", newPerson, value)
	} else {
		fmt.Println("Not found")
	}
}
