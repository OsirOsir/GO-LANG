package main

import "fmt"

func main() {
	myInvoice := newInvoice("Hot Cafe..... \nPhilipo's Bill")

	fmt.Println(myInvoice.format())
}
