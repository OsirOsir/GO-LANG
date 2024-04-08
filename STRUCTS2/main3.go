package main

import "fmt"

func main() {
	myInvoice := newInvoice("Philips Bill")

	myInvoice.updateItems("Chapati", 36.5)
	myInvoice.updateItems("Chai", 10.2)
	myInvoice.updateItems("Beans", 30.5)
	myInvoice.updateItems("Rice", 40.8)
	myInvoice.updateItems("Beaf", 136.5)
	myInvoice.updateItems("Matumbo", 100.1)
	myInvoice.updateTip(10.4)

	fmt.Println(myInvoice.format())
}
