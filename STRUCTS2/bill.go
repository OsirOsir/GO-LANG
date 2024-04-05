package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   int
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}

	return b
}

func (b bill) format() string {
	fs := "Bill breakdown: \n"

	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs = fs + fmt.Sprintf("%-25v   ...$%v \n", k+":", v)
		total = total + v
	}

	//total
	fs = fs + fmt.Sprintf("%-25v   ...$%0.2f", "total:", total)

	return fs
}
