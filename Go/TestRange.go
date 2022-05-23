package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d =%d\n", i, v)
	}
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	for key, value := range map1 {
		fmt.Printf("key is:%d,value is :%f\n", key, value)
	}

	nums := []int{2, 3, 4}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)

	// range 可以便利unicode 字符串

	for i, c := range "hello" {
		fmt.Println(i, c)
	}
}
