package main

import "fmt"

func main() {
	manu := map[string]int{
		"Chapati": 20,
		"Mandazi": 5,
		"Beans":   30,
		"Chai":    10,
		"Uji":     20,
	}

	fmt.Println(manu)
	fmt.Println(manu["Chapati"])

	for k, v := range manu {
		fmt.Println(k, "-", v)
	}

	manu["Chapati"] = 30
	fmt.Println(manu)

}
