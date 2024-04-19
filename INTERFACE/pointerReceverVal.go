package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a Address) Describe() {
	fmt.Printf("State %s country %s\n", a.state, a.country)
}

func main() {
	var d1 Describer
	p1 := Person{"Philip", 36}
	d1 = p1
	d1.Describe()

	p2 = Person{"Alice", 27}
	d1 = &p2
	d1.Describe()
}
