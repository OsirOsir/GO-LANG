package main

import "fmt"

func main() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}

	// for _, i := range byteSlice {
	// 	fmt.Printf("%x ", i)
	// }
	// fmt.Println()

	str := string(byteSlice)
	fmt.Println(str)
}
