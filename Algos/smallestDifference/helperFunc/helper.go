package helperFunc

import (
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {

	sort.Ints(array1)
	sort.Ints(array2)

	first := 0
	second := 0
	smallestDiff, smallestDiffArray := math.MaxInt32, make([]int, 2)

	for first < len(array1) && second < len(array2) {

		diff := absDiff(array1[first], array2[second])

		if diff < smallestDiff {
			smallestDiff = diff
			smallestDiffArray[0] = array1[first]
			smallestDiffArray[1] = array2[second]
		}

		if array1[first] < array2[second] {
			first++
		} else {
			second++
		}
	}

	return smallestDiffArray
}

func absDiff(i, j int) int {

	if j-i < 0 {
		return -1 * (j - i)
	}

	return j - i
}
