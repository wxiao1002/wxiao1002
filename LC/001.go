package lc

import "fmt"

/**
 两数之和
 数组nums 中找出 和为目标值target 的两个数

 解题思路：
	1. 不能循环2次，使用内存换时间
	2. 创建集合map 存储临时值【target与目标数的差】，每次循环判断当前数是否在map 中，存在则返回
	3. 不存在则目标数与当前数的差放入集合中，v 是当前数坐标，
	4. 因为找到的数 是之前数的差 所以坐标返回是 j,i
*/
func main() {
	test1 := []int{1, 2, 3, 4, 5, 6}
	target := 11
	sum := twoSum(test1, target)
	fmt.Println(sum)
}

func twoSum(nums []int, target int) []int {
	len := len(nums)
	var tempMap = make(map[int]int, len)
	for i, v := range nums {
		j, ok := tempMap[v]
		if ok {
			return []int{j, i}
		} else {
			tempMap[target-v] = i
		}
	}
	return []int{-1, -1}
}

// 网上解法
func twoSumCopy(nums []int, target int) (ans []int) {
	idxMap := map[int]int{}
	for i, num := range nums {
		// 优化点 直接 判断集合存在 差值，存在则返回
		if v, ok := idxMap[target-num]; ok {
			ans = []int{v, i}
			break
		}
		idxMap[num] = i
	}
	return
}
