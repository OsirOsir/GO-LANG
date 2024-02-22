package main

import "fmt"

func main() {
	var a [4]int // int array with length 4
	// long  hand declaration.
	a[0] = 30 // array index starts from Zero
	a[1] = 48 // array index  one
	a[2] = 46 // array index two
	a[3] = 98 // array index three
	// for i, elem := range a {
	// 	fmt.Println(elem, i)
	// }
	fmt.Println(a)
	//fmt.Println(a, len(a))
}
