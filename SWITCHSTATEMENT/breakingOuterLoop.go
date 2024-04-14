package main

import (
	"fmt"
	"math/rand"
)

func main() {
randLoop:
	for {
		switch i := rand.Intn(1000); {

		case i%2 == 0:
			fmt.Printf("Generated even number is %d\n", i)
			break randLoop
		}
	}
	// ch := make(chan int)
	// go func() {
	// 	ch <- 1
	// 	ch <- 2
	// 	ch <- 3
	// 	close(ch)
	// }()
	// for value := range ch {
	// 	fmt.Println(value)
	// }

}
