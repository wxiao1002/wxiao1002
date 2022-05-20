package main

import "fmt"

func main() {

	// 创建切片
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	// 打印子切片 ，从1～4
	printSlice(numbers[1:4])

	// 默认下限度是0 上限是len
	printSlice(numbers[:4])
	printSlice(numbers[1:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	// 子切片
	numbers2 := numbers1[:2]
	printSlice(numbers2)
}
func printSlice(sp []int) {
	fmt.Printf("len :%d,cap:%d,slice:%v\n", len(sp), cap(sp), sp)
}
