package main

import "github.com/01-edu/z01"

func PrintComb() {
	for x := '0'; x <= '7'; x++ {
		for y := x + 1; y <= '8'; y++ {
			for z := y + 1; z <= '9'; z++ {
				z01.PrintRune(x)
				z01.PrintRune(y)
				z01.PrintRune(z)
				if x == '7' && y == '8' && z == '9' {
					break
				} else {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
func main() {
	PrintComb()
}
