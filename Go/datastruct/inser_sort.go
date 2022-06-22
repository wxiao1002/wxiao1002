package datastruct

func insertsort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	for i := 1; i < len(array); i++ {
		back := array[i]
		j := i - 1
		for j >= 0 && back < array[j] {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = back
	}
	return array
}
