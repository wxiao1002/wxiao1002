package lc

import "fmt"

func main() {
	var a []int = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(a))
}

func trap(height []int) int {
	res := 0
	for i := 1; i < len(height)-1; i++ {
		left, right := 0, 0
		for j := i; j < len(height); j++ {
			right = getMax(height[j], right)
		}
		for j := i; j > 0; j-- {
			left = getMax(height[j], left)
		}
		res += getMin(left, right) - height[i]
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
