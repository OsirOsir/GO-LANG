package main

func sayGreetings(n string) {
	fmt.printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.printf("Goodbye  %v \n", n)
}

func main() {
	sayGreetings("Philip")
	sayBye("Kyle")
}
