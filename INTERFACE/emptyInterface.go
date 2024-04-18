// package main

// import "fmt"

// type i interface{}

// func found(val i) {

// 	fmt.Printf("Type = %T, value = %v\n", val, val)
// }

// func main() {
// 	s := "Simplilearn"
// 	found(s)
// }

package main

import "fmt"

var i interface{}

func found(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Simplelearn"

	found(s)
	i := 07
	found(i)
}
