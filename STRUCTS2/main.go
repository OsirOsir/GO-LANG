package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")
	mybill.addItem("Chapati", 34.2)
	mybill.addItem("Rice", 54.5)
	mybill.addItem("Beans", 20.8)
	mybill.addItem("Tea", 10.1)
	mybill.updateTip(13.2)
	fmt.Println(mybill.format())

}
