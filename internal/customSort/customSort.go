package customSort

func Bubble(slice []int) ([]int, int) {
	var iterations = 0

	length := len(slice)
	for i := 0; i < length; i++ {
		wasSwapped := false
		for j := 0; j < length-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				wasSwapped = true
			}
			iterations++
		}
		if !wasSwapped {
			break
		}
	}

	return slice, iterations
}
