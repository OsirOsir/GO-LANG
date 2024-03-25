package main

func main() {
	b := [...]int{109, 110, 111}
	p := &b
	p++
}

// Gives an error since In Go, while you can use pointers to arrays and increment or decrement the pointer itself, you cannot perform array
// arithmetic directly using pointers as you would in C or C++.
