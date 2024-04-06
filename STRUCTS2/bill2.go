package main

import "fmt"

type invoice struct {
	name  string
	items map[string]float64
	tip   float64
}

func newInvoice(name string) invoice {
	i := invoice{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return i
}

func (i *invoice) format() string {
	fs := "Summary of Invoice \n"

	var total float64

	for k, v := range i.items {
		fs = fs + fmt.Sprintf("%-25s   ...$%.2f\n", k+":", v)
		total = total + v
	}

	fs = fs + fmt.Sprintf("%-25s   ...$%.2f\n", "Tip:", i.tip)

	fs = fs + fmt.Sprintf("%-25s   ...$%.2f", "Total:", total+i.tip)

	return fs
}

// update tip

func (i *invoice) updateTip(tip float64) {
	i.tip = tip
}

// update Items
func (i *invoice) updateItems(name string, price float64) {
	i.items[name] = price
}
