// First we define an Interface called 'Animal' with a method called 'MakeSound'
package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct{}

func (d Dog) MakeSound() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) MakeSound() string {
	return "Meow!"
}

type Cow struct{}

func (c Cow) MakeSound() string {
	return "Moo!"
}

func LetAnimalSpeak(a Animal) {
	fmt.Println(a.MakeSound())
}

func main() {
	dog := Dog{}
	cat := Cat{}
	cow := Cow{}

	LetAnimalSpeak(dog)
	LetAnimalSpeak(cat)
	LetAnimalSpeak(cow)
}
