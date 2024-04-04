package main

import "fmt"

func main() {
	mybill := newBill("Philip's bill")

	fmt.Println(mybill.format())
}
