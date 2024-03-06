package main

import "fmt"

func main() {
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// The program below uses nested for loops to print the sequence.
// The variable n in line no. 8 stores the number of lines in the sequence.
// In our case it’s 5. The outer for loop iterates i from 0 to 4 and the inner for loop iterates j from 0 to the current value of i.
// The inner loop prints * for each iteration and the outer loop prints a new line at the end of each iteration.
// Run this program and you see the sequence printed as output

// Of course, let's break it down step by step with examples:

// 1. **When i = 0 (First Row):**
//    - In the outer loop, `i` is initially set to 0.
//    - The inner loop iterates from `j = 0` to `j = i`.
//    - So, the inner loop runs only once because when `j = 0`, it satisfies the condition `j <= i` (0 <= 0).
//    - It prints one asterisk for the first row.

//    ```
//    *
//    ```

// 2. **When i = 1 (Second Row):**
//    - In the outer loop, `i` is incremented to 1.
//    - The inner loop iterates from `j = 0` to `j = i`.
//    - For `i = 1`, the inner loop runs twice because it satisfies the condition `j <= i` for both `j = 0` and `j = 1`.
//    - It prints two asterisks for the second row.

//    ```
//    **
//    ```

// 3. **When i = 2 (Third Row):**
//    - In the outer loop, `i` is incremented to 2.
//    - The inner loop iterates from `j = 0` to `j = i`.
//    - For `i = 2`, the inner loop runs three times because it satisfies the condition `j <= i` for `j = 0`, `j = 1`, and `j = 2`.
//    - It prints three asterisks for the third row.

//    ```
//    ***
//    ```

// 4. **When i = 3 (Fourth Row):**
//    - In the outer loop, `i` is incremented to 3.
//    - The inner loop iterates from `j = 0` to `j = i`.
//    - For `i = 3`, the inner loop runs four times because it satisfies the condition `j <= i` for `j = 0`, `j = 1`, `j = 2`, and `j = 3`.
//    - It prints four asterisks for the fourth row.

//    ```
//    ****
//    ```

// 5. **When i = 4 (Fifth Row):**
//    - In the outer loop, `i` is incremented to 4.
//    - The inner loop iterates from `j = 0` to `j = i`.
//    - For `i = 4`, the inner loop runs five times because it satisfies the condition `j <= i` for `j = 0`, `j = 1`, `j = 2`, `j = 3`, and `j = 4`.
//    - It prints five asterisks for the fifth row.

//    ```
//    *****
//    ```

// So, the pattern emerges where each row has an increasing number of asterisks, starting from one and going up to five. This is achieved by carefully controlling the iteration of the inner loop based on the value of `i`.
