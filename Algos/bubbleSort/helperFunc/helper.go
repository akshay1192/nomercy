package helperFunc

func BubbleSort(array []int ) []int{

	for i := 0; i < len(array)-1; i++ {

		swappedOnce := false
		for j := 0; j < len(array)-i-1; j++{
			if array[j] > array[j+1] {

				// no need of temp var for swapping
				array[j],array[j+1] = array[j+1],array[j]
				swappedOnce = true
			}
		}

		if !swappedOnce {
			break
		}
	}
	return array
}