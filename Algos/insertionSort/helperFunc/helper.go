package helperFunc

func InsertionSort(array []int) []int {

	for i := 0; i < len(array)-1; i++ {

		swappedOnce := false
		for j := i + 1; j > 0; j-- {

			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
				swappedOnce = true
			}

			if !swappedOnce {
				break
			}
		}
	}

	return array

}
