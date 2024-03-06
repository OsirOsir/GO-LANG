package main

import (
	"fmt"
	"math"
)

func maleName(s string) {
	fmt.Printf("%s is a male name\n", s)
}

func femaleName(s string) {
	fmt.Printf("%s is a Female name\n", s)
}

func peopleNames(s []string, f func(string)) {
	for _, v := range s {
		f(v)
	}

}

func areaCircle(n float64) float64 {
	return math.Pi * n * n
}

func main() {
	// maleName("Philip")
	// femaleName("Iddah")

	peopleNames([]string{"Philip", "Andrew", "Fredrick", "Joshua"}, maleName)
	peopleNames([]string{"Alice", "Yuanita", "Eucabeth", "Carolyn"}, femaleName)

	a1 := areaCircle(15.5)
	a2 := areaCircle(34.5)
	fmt.Printf("area of Circle 1 and Circle 2 is %.1f and %.1f respectively \n", a1, a2)
}
