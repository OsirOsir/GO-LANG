package main

import "fmt"

func PrintComb() {
	for x := '0'; x <= '7'; x++ {
		for y := x + 1; y <= '8'; y++ {
			for z := y + 1; z <= '9'; z++ {
				fmt.Print(string(x))
				fmt.Print(string(y))
				fmt.Print(string(z))
				// if x == '7' && y == '8' && z == '9' {
				// 	break
				// } else {
				fmt.Print(",")
				fmt.Print(" ")
				// }
			}
		}
	}
	fmt.Println()
}
func main() {
	PrintComb()
}
