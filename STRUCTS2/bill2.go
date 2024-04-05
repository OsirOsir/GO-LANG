package main

import (
	"fmt"
)

type invoice struct {
	name  string
	items map[string]float64
	tip   int
}

func newInvoice(name string) invoice {
	i := invoice{
		name:  name,
		items: map[string]float64{"Chapati": 30.5, "Beans": 60.3},
		tip:   0,
	}

	return i
}

func (i invoice) format() string {
	fs := "Invoice summary\n"

	var total float64 = 0

	for k, v := range i.items {
		fs += fmt.Sprintf("%-25s  ...$%.2f\n", k, v)
		total += v
	}

	fs = fs + fmt.Sprintf("%-25s  ...$%.2f", "Total:", total)
	return fs
}
