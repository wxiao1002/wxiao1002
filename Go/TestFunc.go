package main

import (
	"fmt"
	"math"
)

func main() {
	c := max(1, 2)
	fmt.Println(c)
	var s1, s2 string = swap("123", "321")
	fmt.Println(s1, s2)

	// 声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	// 使用函数
	fmt.Println(getSquareRoot(9))

	nextNumber := getSequence()
	fmt.Println(nextNumber())

	var c1 Circle
	c1.radius = 20.00
	fmt.Println(c1.getArea())
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
func swap(x, y string) (string, string) {
	return y, x
}

// 定义结构体
type Circle struct {
	radius float64
}

// 该方法数据 Circle 结构体中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
