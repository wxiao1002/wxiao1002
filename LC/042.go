package lc

import "fmt"

func main() {
	var a []int = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(a))
}

func trap(height []int) int {
	res, n := 0, len(height)
	for i := 1; i < n-1; i++ {
		leftMax, rightMax := 0, 0
		for j := i; j < n; j++ {
			rightMax = getMax(height[j], rightMax)
		}
		for j := i; j > 0; j-- {
			leftMax = getMax(height[j], leftMax)
		}
		res += getMin(leftMax, rightMax) - height[i]
	}
	return res
}

func getMax(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
func getMin(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

func trap2(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax, res := 0, 0, 0
	for left <= right {
		leftMax = getMax(leftMax, height[left])
		rightMax = getMax(rightMax, height[right])

		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}
