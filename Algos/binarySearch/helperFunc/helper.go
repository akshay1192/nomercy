package helperFunc

func BinarySearchRecursion(array []int, target int) int {
	return binarySearchRecursionHelper(array, target, 0, len(array)-1)
}

func binarySearchRecursionHelper(array []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := right - (right-left)/2
	if array[mid] == target {
		return mid
	}

	if array[mid] > target {
		return binarySearchRecursionHelper(array, target, left, mid-1)
	}

	return binarySearchRecursionHelper(array, target, mid+1, right)

}

func BinarySearchIterative(array []int,target int) int {
	right := len(array)-1
	left := 0

	for left <= right {
		mid := right - (right - left)/2

		if array[mid] == target {
			return mid
		} else if array[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return -1
}