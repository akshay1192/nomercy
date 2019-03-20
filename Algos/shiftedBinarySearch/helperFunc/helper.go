package helperFunc

func ShiftedBinarySearch(array []int, target int) int {

	pivot := findPivot(array)
	len := len(array)

	if pivot == -1 {
		return binarySearch(array, target, 0, len-1)
	}

	if array[0] <= target && array[pivot] >= target {
		return binarySearch(array, 0, pivot, target)
	} else {
		return binarySearch(array, pivot+1, len-1, target)
	}

	return -1

}

func findPivot(array []int) int {

	pivot, low, len := -1, 0, len(array)
	high := len - 1

	for low <= high {

		mid := low + (high-low)/2

		if array[mid] > array[mid+1] {
			pivot = mid
			break
		} else if array[mid] < array[mid-1] {
			pivot = mid - 1
			break
		}

		if array[low] > array[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}

	return pivot
}

func binarySearch(array []int, left, right, target int) int {

	index, low, high := -1, left, right

	for low <= high {
		mid := low + (high-low)/2

		if array[mid] == target {
			return mid
		} else if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return index
}
