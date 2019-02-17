package helperFunc

import (
	"sort"
)

func CheckSum(inputArr []int, targetSum int) []int {

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
