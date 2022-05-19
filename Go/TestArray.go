package main

import "fmt"

func main() {

	var balance = [5]float32{100.0, 1.0, 3.4, 32.0, 45.1}
	fmt.Println(&balance)

	cc := [...]int{1, 2, 3}
	fmt.Println(cc[0])
	// 创建数组
	values := [][]int{}
	// 创建1维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	// 向二维数组中添加一维数组
	values = append(values, row1)
	values = append(values, row2)
	fmt.Println(values[1][1])

	sites := [2][2]string{}
	sites[0][0] = "hello"
	fmt.Println(sites)

	for i := range sites {
		fmt.Println(sites[i])
	}
}
