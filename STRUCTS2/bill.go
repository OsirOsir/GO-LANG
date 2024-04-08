package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
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
	//tip
	fs = fs + fmt.Sprintf("%-25v   ...$%0.2f\n", "Tip:", b.tip)

	//total
	fs = fs + fmt.Sprintf("%-25v   ...$%0.2f", "Total:", total+b.tip)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b bill) addItem(name string, price float64) {
	b.items[name] = price
}
