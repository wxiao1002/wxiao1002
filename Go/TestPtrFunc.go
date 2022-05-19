package main

import "fmt"

func main() {
	x := 100
	y := 200
	swapVar(&x, &y)
	fmt.Printf("x is: %d,y is %d", x, y)
}

func swapVar(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
