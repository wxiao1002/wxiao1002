package datastruct

func Sort(arr []int, low, high int) {

	flag := partition(arr, low, high)

	Sort(arr, low, flag)
	Sort(arr, flag+1, high)

}

func partition(arr []int, low, high int) int {

	provite := arr[low]

	for low <= high {

		for arr[high] > provite {
			high--
		}
		arr[low] = arr[high]

		for arr[low] < provite {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = provite
	return low

}
