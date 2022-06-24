package datastruct


func sort(array []int) []int {
	count := len(array)
	for i := 0; i < count; i++ {
		for j := 0; j < count-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

