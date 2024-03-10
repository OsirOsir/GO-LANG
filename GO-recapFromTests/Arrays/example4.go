package main

import "fmt"

func main() {
	a := [...]int{3}
	var b [3]string
	b = [3]string{"Philip"}
	fmt.Println(a, len(a))
	fmt.Println(b, len(b))
}
