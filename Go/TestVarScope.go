package main

import "fmt"

// 全局变量
var g int

func main() {
	// 局部变量
	var a, b int
	a = 10
	b = 20
	g = a + b
	fmt.Println(a, b, g)
}
