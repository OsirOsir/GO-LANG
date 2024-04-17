package main

import "fmt"

type Animal interface {
	AnimalSound() string
}

type Dog struct{}
type Cat struct{}
type Cow struct{}

func (d Dog) AnimalSound() string {
	return "Woof!"
}

func (c Cat) AnimalSound() string {
	return "Meow!"
}

func (c Cow) AnimalSound() string {
	return "Moo!"
}

func listenToSounds(a Animal) {
	fmt.Println(a.AnimalSound())
}

func main() {
	d := Dog{}
	c := Cat{}
	co := Cow{}

	listenToSounds(d)
	listenToSounds(c)
	listenToSounds(co)

}
