package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	PrintSlice(numbers)

	var number1 []int
	if number1 == nil {
		fmt.Println("is nil")
	}
}
func PrintSlice(sp []int) {
	fmt.Printf("len :%d,cap:%d,slice:%v\n", len(sp), cap(sp), sp)
}
