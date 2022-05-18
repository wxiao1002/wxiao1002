package main

import "fmt"

func main() {
	var a = 60         // 0011 1100
	var b = 13         // 0000 1101
	fmt.Println(a & b) // 0000 1100
	fmt.Println(a | b) // 0011 1101
	fmt.Println(a ^ b) // 0011 0001
}
