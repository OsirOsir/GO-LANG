package main

import "fmt"

type invoice struct {
	name  string
	items map[string]float64
	tip   int
}

func newInvoice(name string) invoice {
	i := invoice{
		name:  name,
		items: map[string]float64{"Chapati": 30.5, "Beans": 60.8},
		tip:   0,
	}
	return i
}

func (i invoice) format() string {
	fs := "Summary of Invoice \n"

	var total float64

	for k, v := range i.items {
		fs = fs + fmt.Sprintf("%-25s   ...$%.2f\n", k+":", v)
		total = total + v
	}

	fs = fs + fmt.Sprintf("%-25s   ...$%.2f", "Total:", total)

	return fs
}
