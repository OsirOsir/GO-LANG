package main

import "fmt"

func assertPanicSort(i interface{}) {
	v, ok := i.(int)

	fmt.Println(v, ok)
}

func main() {
	var n interface{} = 56
	assertPanicSort(n)

	var i interface{} = "Philip Osir"

	assertPanicSort(i)
}
