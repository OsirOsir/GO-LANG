package main

import "fmt"

type invoice struct {
	name string
	item map[string]float64
	tip  float64
}

func newInvoice(name string) invoice {
	i := invoice{
		name: name,
		item: map[string]float64{},
		tip:  0,
	}

	return i
}

func (i invoice) format() string {
	fs := "Invoice Summary\n"

	total := 0.0

	for k, v := range i.item {
		fs += fmt.Sprintf("%-25s   ...%.1f\n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25s   ...%.1f\n", "Tip:", i.tip)
	fs += fmt.Sprintf("%-25s   ...%.1f\n", "Total:", total+i.tip)

	return fs
}

func (i invoice) updateItem(item string, price float64) {
	i.item[item] = price
}

func (i *invoice) updateTip(tip float64) {
	i.tip = tip
}
