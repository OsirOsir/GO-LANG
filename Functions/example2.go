package main

import "fmt"

func sayGreetings(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye  %v \n", n)
}

func cycleArg(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}
func main() {
	// sayGreetings("Philip")
	// sayBye("Kyle")

	cycleArg([]string{"Joshua", "Opondo", "Kenneth", "Eucabeth"}, sayGreetings)
	cycleArg([]string{"Joshua", "Opondo", "Kenneth", "Eucabeth"}, sayBye)
}
