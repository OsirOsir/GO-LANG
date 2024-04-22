package main

import "fmt"

func main() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}

	fmt.Printf("the bytes are %x: ", byteSlice)

	str := string(byteSlice)
	fmt.Println(str)
}
