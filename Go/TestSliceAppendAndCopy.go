package main

import "fmt"

func main() {
	var numbers []int
	PrintSlice1(numbers)

	// 追加元素
	numbers = append(numbers, 0)
	numbers = append(numbers, 1)
	PrintSlice1(numbers)
	// 追加多个元素
	numbers = append(numbers, 2, 3, 4)
	PrintSlice1(numbers)

	numbers1 := make([]int, len(numbers), cap(numbers)*2)

	// 拷贝之前的内容到 新的
	copy(numbers1, numbers)
	PrintSlice1(numbers1)
}
func PrintSlice1(sp []int) {
	fmt.Printf("len :%d,cap:%d,slice:%v\n", len(sp), cap(sp), sp)
}
