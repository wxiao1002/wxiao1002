package main

import "fmt"

func main() {
	_, num, str := numbers()
	fmt.Println(num, str)
}

func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
