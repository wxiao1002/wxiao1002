package datastruct

func main() {

}

func search1(array []int, target int) int {
	low, high := 0, len(array)
	mid := low + (high-low)/2
	for low <= high {
		if target > array[mid] {
			low = mid + 1
		}
		if target < array[mid] {
			high = mid - 1
		}
		return mid
	}
	return -1
}

func searh2(array []int, low, high, target int) int {

	if low > high {
		return -1
	}
	mid := low + (high-low)>>1
	if target > array[mid] {
		return searh2(array, mid+1, high, target)
	} else if target < array[mid] {
		return searh2(array, low, mid-1, target)
	} else {
		return mid
	}
}

// 查找第一个等于目标值的元素

func search3(array []int, target int) int {
	low, high := 0, len(array)
	for low <= high {
		mid := low + (high-low)>>1
		if array[mid] > target {
			high = mid - 1
		} else if array[mid] < high {
			low = mid + 1
		} else {
			if mid == 0 || array[mid-1] != target {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// 第一个大于等于
func search4(array []int, target int) int {
	low, high := 0, len(array)
	for low <= high {
		mid := low + (high-low)>>1
		if array[mid] >= target {
			if mid == 0 || array[mid-1] < target {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
