package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Print(i)
	}
	fmt.Println()
}

// another way to write  this

// package main

// import "fmt"

// func main() {
// 	i := 0
// 	for i <= 10 {
// 		fmt.Printf("%d ", i)
//         i += 2
// 	}
// }
