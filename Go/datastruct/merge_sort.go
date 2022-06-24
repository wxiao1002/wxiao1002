package datastruct

func MergeSort(arr []int, left, right int) {

	if left < right {
		mid := (left + right) / 2
		MergeSort(arr, left, mid)
		MergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}
func merge(arr []int, left, mid, rigtht int) {
	temp := make([]int, rigtht-left+1)

	i, j, index := left, mid+1, 0

	for i < mid && j < rigtht {
		if arr[i] < arr[j] {
			temp[index] = arr[i]
			i++
		} else {
			temp[index] = arr[j]
			j++
		}
		index++
	}

	for i < mid {
		temp[index] = arr[i]
		index++
		i++
	}

	for j < rigtht {
		temp[index] = arr[j]
		index++
		j++
	}

	for i := left; i <= rigtht; i++ {
		arr[i] = temp[i]
	}

}
