package customSort

func Bubble(slice []int) ([]int, int) {
	var iterations = 0
	var localSlice = make([]int, len(slice))
	copy(localSlice, slice)

	length := len(localSlice)
	for i := 0; i < length; i++ {
		wasSwapped := false
		for j := 0; j < length-1-i; j++ {
			if localSlice[j] > localSlice[j+1] {
				localSlice[j], localSlice[j+1] = localSlice[j+1], localSlice[j]
				wasSwapped = true
			}
			iterations++
		}
		if !wasSwapped {
			break
		}
	}

	return localSlice, iterations
}

func Shaker(slice []int) ([]int, int) {
	var iterations = 0
	var left = 0
	var right = len(slice) - 1
	var control = right

	var localSlice = make([]int, len(slice))
	copy(localSlice, slice)

	for left < right {
		for i := 0; i < right; i++ {
			if localSlice[i] > localSlice[i+1] {
				localSlice[i], localSlice[i+1] = localSlice[i+1], localSlice[i]
				iterations++
				control = i
			}
		}
		right = control

		for j := right; j > left; j-- {
			if localSlice[j] < localSlice[j-1] {
				localSlice[j], localSlice[j-1] = localSlice[j-1], localSlice[j]
				iterations++
				control = j
			}
		}
		left = control
	}

	return localSlice, iterations
}

func Insertion(slice []int) ([]int, int) {
	var iterations = 0
	var localSlice = make([]int, len(slice))
	copy(localSlice, slice)

	length := len(localSlice)

	for i := 0; i < length; i++ {
		minIndex := i
		for j := i; j < length; j++ {
			if localSlice[j] < localSlice[minIndex] {
				minIndex = j
			}
			iterations++
		}
		if localSlice[i] != localSlice[minIndex] {
			localSlice[i], localSlice[minIndex] = localSlice[minIndex], localSlice[i]
		}
	}

	return localSlice, iterations
}
