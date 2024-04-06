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
	st := "Invoice Summary\n"

	var total float64 = 0

	for k, v := range i.item {
		st += fmt.Sprintf("%-25s   ...%.1f\n", k+":", v)
		total += v
	}

	st += fmt.Sprintf("%-25s   ...%.1f\n", "Tip:", i.tip)
	st += fmt.Sprintf("%-25s   ...%.1f", "Total:", total+i.tip)

	return st
}

func (i invoice) updateItem(item string, price float64) {
	i.item[item] = price
}

func (i *invoice) updateTip(tip float64) {
	i.tip = tip
}
