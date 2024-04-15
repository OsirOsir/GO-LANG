package main

import (
	"fmt"
	"os"
)

type Invoice struct {
	name string
	item map[string]float64
	tip  float64
}

func newInvoice(name string) Invoice {
	i := Invoice{
		name: name,
		item: map[string]float64{},
		tip:  0,
	}
	return i
}

func (i Invoice) format() string {
	fs := "Invoice Summary\n"

	total := 0.0

	for k, v := range i.item {
		fs = fs + fmt.Sprintf("%-25s   ...$%.1f\n", k+":", v)
		total = total + v
	}
	fs = fs + fmt.Sprintf("%-25s   ...$%.1f\n", "Tip:", i.tip)

	fs = fs + fmt.Sprintf("%-25s   ...$%.1f", "Total:", total+i.tip)
	return fs
}

func (i Invoice) updateItems(itemName string, price float64) {
	i.item[itemName] = price
}

func (i *Invoice) updateTip(tipValue float64) {
	i.tip = tipValue
}

// save file

func (i Invoice) save() {
	data := []byte(i.format())

	err := os.WriteFile("bill3/"+i.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Invoice saved to File")
}
