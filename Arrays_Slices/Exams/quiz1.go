// 1. **Initialization and Access:**
//   - Declare an array named `numbers` containing integers from 1 to 5.
//   - Print the third element of the array.
package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5}

	fmt.Println("The third number is :", numbers[2])
}
