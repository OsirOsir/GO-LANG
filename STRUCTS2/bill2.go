package main

import (
	"fmt"
	"os"
)

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
		st += fmt.Sprintf("%-25s   ...%.2f\n", k+":", v)
		total += v
	}
	st += fmt.Sprintf("%-25s   ...%.2f\n", "Tip:", i.tip)
	st += fmt.Sprintf("%-25s   ...%.2f", "Total:", total+i.tip)

	return st
}

func (i invoice) updateItem(item string, price float64) {
	i.item[item] = price
}

func (i *invoice) updateTip(tip float64) {
	i.tip = tip
}

// save invoice

func (i invoice) save() {
	data := []byte(i.format())

	err := os.WriteFile("bills2/"+i.name+".txt", data, 0644)
	if err != nil {
		panic(err) // away of stoping the flow if there is an errror
	}
	fmt.Println("Invoice was saved to file")
}
