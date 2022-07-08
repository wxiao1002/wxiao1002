package main

//有n个筹码。第 i 个筹码的位置是position[i]。
//
//我们需要把所有筹码移到同一个位置。在一步中，我们可以将第 i 个筹码的位置从position[i]改变为:
//
//position[i] + 2或position[i] - 2，此时cost = 0
//position[i] + 1或position[i] - 1，此时cost = 1
//返回将所有筹码移动到同一位置上所需要的 最小代价 。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/minimum-cost-to-move-chips-to-the-same-position
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

	// 这道题定位是简单题，但是中文翻译后感觉是文字游戏

	a := []int{1, 2, 3}
	minCostToMoveChips(a)
}

func minCostToMoveChips(position []int) int {
	// 假设我们将所有筹码都移动到最后一个上面
	// 从倒数第二移动 到  最后一个位置需要代价是 1
	// 从倒数第三个移 到 最后一个位置的代价是 0 因为 差值是2
	// 从倒数第四个移动 到 最后一个 代价 1 因为先移动到倒数第二 代价是0 再移动到倒数第一 代价是1

	//

	rk, ri := 0, 0

	for _, v := range position {
		if v%2 == 0 {
			rk += 1
		} else {
			ri += 1
		}
	}
	return getMin(rk, ri)

}
