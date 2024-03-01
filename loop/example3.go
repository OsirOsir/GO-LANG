// continue

package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%v,", i)
	}
	fmt.Println()
}

// In the above program the line if i%2 == 0 checks if the remainder of dividing i by 2 is 0.
// If it is zero, then the number is even and continue statement is executed and the control moves to the next iteration of the loop.
// Hence the print statement after the continue will not be called and the loop proceeds to the next iteration.

// The output of the above program is
// 2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,32,34,36,38,40,42,44,46,48,50,
