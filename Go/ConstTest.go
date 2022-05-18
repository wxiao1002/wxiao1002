package main

import (
	"fmt"
	"unsafe"
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {

	const LENGTH = 10
	const WIDTH int = 20
	//const a, b, c = 1, false, "ss"
	//
	//area := LENGTH * WIDTH
	//fmt.Println(area)
	//println()
	//println(a, b, c)

	fmt.Println(a, b, c)
	test_iota()

}

const (
	i1 = 1 << iota
	j1 = 3 << iota
	k1
	l1
)

func test_iota() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	fmt.Println(i1, j1, k1, l1)

}
