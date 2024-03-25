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

}
