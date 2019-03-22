package helperFunc

import "fmt"


// we are also calculating the subarray having max sum
func KadanesAlgorithm(array []int) int {

	maxSoFar,subArray,maxEndingHere := array[0],make([]int,2),array[0]
	start,end := 0,0


	for i := 1; i < len(array); i++ {

		maxEndingHere = max(maxEndingHere,array[i])

		if maxEndingHere <= 0 {
			start = i
			end = i
		}

		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
			end = i
			subArray = []int{start,end}
		}
	}

	// exclude start index from subarray if array[start] was negative
	if array[start] < 0 && start != end {
		subArray = []int{start+1,end}
	}

	fmt.Println("Subarray : ",array[subArray[0]:subArray[1]+1])
	return maxSoFar
}


func max(i,j int)int {

	if i+j > j {
		return i+j
	}

	return j
}
