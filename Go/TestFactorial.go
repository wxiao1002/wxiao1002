package main

import "fmt"

func main() {
	fmt.Println(Factrorial(10))
}

func Factrorial(n uint64) uint64 {
	if n > 0 {
		return n * Factrorial(n-1)
	}
	return 1
}
