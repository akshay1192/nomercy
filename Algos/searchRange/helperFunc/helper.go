package helperFunc

func SearchForRange(array []int, target int) []int {
	leftIndex := binarySearchHelper(array, target, true)
	rightIndex := binarySearchHelper(array, target, false)

	return []int{leftIndex, rightIndex}
}

func binarySearchHelper(array []int, target int, left bool) int {
	len, low, index := len(array), 0, -1
	high := len - 1

	for low <= high {
		mid := low + (high-low)/2

		switch {
		case array[mid] == target:
			{
				index = mid

				if left {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}

		case array[mid] < target:
			low = mid + 1

		default:
			high = mid - 1

		}
	}
	return index
}

func SearchForRangeIfElse(array []int, target int) []int {
	leftIndex := binarySearchHelperIfElse(array, target, true)
	rightIndex := binarySearchHelperIfElse(array, target, false)

	return []int{leftIndex, rightIndex}
}

func binarySearchHelperIfElse(array []int, target int, left bool) int {

	len := len(array)

	low := 0
	high := len - 1

	index := -1

	for low <= high {
		mid := low + (high-low)/2

		if array[mid] == target {
			index = mid

			if left {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if array[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return index

}
