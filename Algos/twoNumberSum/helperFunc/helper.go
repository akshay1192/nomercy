package helperFunc

import (
	"sort"
)

func CheckSumLoop(inputArr []int, targetSum int) []int {
	for i := 0; i < len(inputArr)-1; i++ {
		for j := i + 1; j < len(inputArr); j++ {
			if inputArr[i]+inputArr[j] == targetSum {
				result := []int{inputArr[i], inputArr[j]}
				sort.Ints(result)
				return result
			}
		}
	}
	return []int{}
}

func CheckSumSort(inputArr []int, targetSum int) []int {

	sort.Ints(inputArr)
	left := 0
	right := len(inputArr) - 1

	for left < right {

		switch {

		case inputArr[left]+inputArr[right] == targetSum:
			return []int{inputArr[left], inputArr[right]}
		case inputArr[left]+inputArr[right] < targetSum:
			left++
		case inputArr[left]+inputArr[right] > targetSum:
			right--
		}

	}

	return []int{}

}

func CheckSumHash(inputArr []int, targetSum int) []int {

	hashMap := map[int]bool{}

	for i := 0; i < len(inputArr); i++ {

		requiredNum := targetSum - inputArr[i]

		if hashMap[requiredNum] {
			result := []int{requiredNum, inputArr[i]}
			sort.Ints(result)
			return result
		}

		hashMap[inputArr[i]] = true
	}

	return []int{}
}
