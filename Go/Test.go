package main

import "fmt"

var (
	ab int
	cd string
)

func main() {
	var a = "Test"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d int
	d = 12
	fmt.Println(d)

	var e bool
	fmt.Println(e)

	var c1 *int
	fmt.Println(c1)

	d2 := "hello"
	fmt.Println(d2)

	var i, j = 1, 2
	fmt.Println(i, j)

	s, u := 1, "23"
	fmt.Println(s, u)

	fmt.Println(ab, cd)
}
