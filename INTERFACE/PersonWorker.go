package main

import "fmt"

type Worker interface {
	work()
}

type Person struct {
	name string
}

func (p Person) work() {
	fmt.Println(p.name, "is working")
}

func describe(w Worker) {
	fmt.Printf("Interface type %T value %s\n", w, w)
}

func main() {
	p := Person{
		name: "Philip",
	}

	var w Worker = p
	describe(w)
	p.work()
}
