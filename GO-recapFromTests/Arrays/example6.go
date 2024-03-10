package main

import "fmt"

func main() {
	a := [...]string{"Philip", "Henril", "Joshua", "Otieno"}

	b := a
	b[1] = "Okeyo"
	b[2] = "Adhiambo"

	fmt.Println(b)
	fmt.Println(a)
}
