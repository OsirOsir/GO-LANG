package main

import "fmt"

func like(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I learn at %s\n", i.(string))
	case int:
		fmt.Printf("And I am %d years old\n ", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	like("Tech school")
	like(65)
	like(6.2)

}
