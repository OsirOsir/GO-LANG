// breaks

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated the moment its i > 5
		}
		fmt.Print(i)
	}
	fmt.Println()
}

// In the above program, the value of i is checked during each iteration.
// If i is greater than 5 then break executes and the loop is terminated.
// The print statement just after the for loop is then executed.

// The above program will output,
// 1 2 3 4 5
// line after for loop
