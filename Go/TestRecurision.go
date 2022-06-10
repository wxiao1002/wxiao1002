package go_test

func f(n int) int {
	if n == 1 {
		return 1
	}
	return f(n-1) + 1
}
func f2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return f(n-1) + f(n-2)
}
