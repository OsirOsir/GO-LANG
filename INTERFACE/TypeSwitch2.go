package main

import "fmt"

func assert(i interface{}) {

	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf(" Unkown type\n")
	}
}

func main() {
	name := "Philip"
	Age := 34
	weight := 4.5

	assert(name)
	assert(Age)
	assert(weight)

}
