package helperFunc

import "math"

func FindThreeLargestNumbersBubbleSort(array []int) []int {

	for i := 0; i < 3; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}

	return array[len(array)-3:]
}

func FindThreeLargestNumbers(array []int) []int {
	result := [3]int{math.MinInt32, math.MinInt32, math.MinInt32}

	for _, val := range array {
		switch {
		case val > result[2]:
			result[0] = result[1]
			result[1] = result[2]
			result[2] = val

		case val > result[1]:
			result[0] = result[1]
			result[1] = val

		case val > result[0]:
			result[0] = val
		}
	}

	return result[0:3]
}
