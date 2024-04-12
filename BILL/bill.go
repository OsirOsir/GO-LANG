package main

import "fmt"

type bill struct {
	name string
	item map[string]float64
	tip  float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		item: map[string]float64{},
		tip:  0,
	}
	return b
}

func (b bill) format() string {
	fs := "Bills Summary\n"
	total := 0.0

	for k, v := range b.item {
		fs += fmt.Sprintf("%-25s   ...%.2f\n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-25s   ...%.2f\n", "Tip:", b.tip)
	fs += fmt.Sprintf("%-25s   ...%.2f", "Total:", total+b.tip)

	return fs
}

func (b bill) updateItem(item string, price float64) {
	b.item[item] = price
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}
