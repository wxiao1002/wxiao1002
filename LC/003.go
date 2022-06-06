package lc

// 双指针 
func lengthOfLongestSubstring(s string) int {
	
	byteMap := map[byte]int {}

	length := len(s)

	begin, ans := -1,0

	for i := 0; i < length; i++ {
		if i!= 0 {
			delete(byteMap,s[i-1])
		}
		for ans +1 < length && byteMap[s[begin+1]]==0  {
			byteMap[s[begin+1]]++
            begin++
		}
		ans = max(ans, begin - i + 1)
	}
	return ans

}
func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}