package helperFunc

import (
	"sort"
)

func ThreeNumberSum(array []int, target int) [][]int {

	var results = make([][]int, 0)
	sort.Ints(array)

	for i := 0; i < len(array)-2; i++ {
		left := i + 1
		right := len(array) - 1

		for left < right {
			if array[left]+array[right]+array[i] == target {
				result := []int{array[i], array[left], array[right]}
				results = append(results, result)
				left++
				right--
			} else if array[left]+array[right]+array[i] < target {
				left++
			} else {
				right--
			}
		}
	}
	return results
}

func ThreeNumberSumNaiveApproach(array []int, target int) [][]int {

	sort.Ints(array)
	result := [][]int{}

	length := len(array)

	for i := 0; i < length - 2; i++ {
		for j := i+1; j < length - 1; j++ {
			for k := j+1; k < length; k++ {
				if array[i]+ array[j] + array[k] == target {
					answer := []int{array[i],array[j],array[k]}
					result = append(result,answer)
				}
			}
		}
	}

	return result
}
