package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("%x", &a)

	var c int = 20 // 声明实际变量
	var ip *int    // 声明指针
	ip = &c        // 指针变量的存储地址
	fmt.Printf("a 变量地址是%x\n", &c)

	fmt.Printf("*ip 变量的值%d\n", *ip)

	arr := [3]int{100, 200, 300}
	var i int
	var ptr [3]*int
	for i = 0; i < 3; i++ {
		ptr[i] = &arr[i]
	}
	for i = 0; i < 3; i++ {
		fmt.Printf("a[%d] is%d", i, *ptr[i])
	}

}
