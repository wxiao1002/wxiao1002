package main

func main() {
	fibonacci(4)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
