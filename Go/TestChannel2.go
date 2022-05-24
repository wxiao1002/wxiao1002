package main

import "fmt"

func main() {
	// 定义一个 存储整数类型的带缓冲区 通道
	ch := make(chan int, 2)
	//  同时发送两个数据 而不需要立刻读取数据
	ch <- 1
	ch <- 2
	// 获取数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
