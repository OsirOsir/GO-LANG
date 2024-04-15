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
		item: map[string]float64{"Chapo": 34, "Chai": 10.2},
		tip:  10,
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
