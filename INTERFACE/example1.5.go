package main

import "fmt"

type Animal interface {
	makeSound() string
}

type Dog struct{}
type Cat struct{}
type Cow struct{}

func (d Dog) makeSound() string {
	return "Woof!"
}

func (c Cat) makeSound() string {
	return "Meau!"
}

func (c Cow) makeSound() string {
	return "Moo!"
}

func AnimalSpeak(a Animal) {
	fmt.Println(a.makeSound())
}

func main() {
	dog := Dog{}
	cat := Cat{}
	cow := Cow{}

	AnimalSpeak(dog)
	AnimalSpeak(cat)
	AnimalSpeak(cow)
}
