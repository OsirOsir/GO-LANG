package main

import "fmt"

func main() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c", i)
	}
	fmt.Println()
}

// package main

// import "fmt"

// func main() {
// 	for i := 'z'; i >= 'a'; i-- {
// 		fmt.Print(string(i))
// 	}
// 	fmt.Println()
// }
