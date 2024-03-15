// package main

// import "fmt"

// func main() {
// 	wages := make(map[string]int)

// 	wages["Alice"] = 20000
// 	wages["Robert"] = 30000
// 	wages["Naomie"] = 10000
// 	fmt.Println("The employee wages are :", wages)
// }

package main

import "fmt"

func main() {
	wages := map[string]int{
		"Alice":  30000,
		"Robert": 20000,
		"Steven": 80000,
	}
	fmt.Println("The wages of employees", wages)
}
