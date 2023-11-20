package merge_sort

func mergeSort(input []int) (output []int) {
	if len(input) < 2 {
		return input
	}

	middle := len(input) / 2
	a := mergeSort(input[:middle])
	b := mergeSort(input[middle:])

	return Merge(a, b)
}

func Merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size)

	// пока не перебрали все в левой или в правой части
	for i < len(left) || j < len(right) {
		// если все перебрали в правой или если лавая часть меньше
		if j == len(right) || i < len(left) && left[i] < right[j] {
			slice[i+j] = left[i]
			i++
		} else {
			// если правая часть меньше или в левой больше нету
			slice[i+j] = right[j]
			j++
		}

	}
	return slice
}
