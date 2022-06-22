package datastruct

func selectsort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	for i := 0; i < len(array)-1; {
		minpos := i
		for j := i; j < len(array); {
			if array[j] < array[i] {
				minpos = j
			}

			array[i], array[minpos] = array[minpos], array[minpos]
		}
		i++
	}
}
