//Count number of occurrences (or frequency) in a sorted array
//Given a sorted array arr[] and a number x, write a function that counts the occurrences of x in arr[]

package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 5, 6}
	targetNum := 4

	// linear search ( O(n) )
	fmt.Println("Linear search : ", linearSearch(array, targetNum))

	// binary search ( O(logN) + count )
	fmt.Println("Binary search : ", binarySearch(array, targetNum))

	// binary search ( O(logN) )
	left := binarySearchOptimised(array, targetNum, true)
	right := binarySearchOptimised(array, targetNum, false)

	if left == 0 && right == 0 {
		fmt.Println("Binary search optimised : ", 0)
	} else {
		fmt.Println("Binary search optimised : ", right-left+1)
	}

}

func linearSearch(array []int, targetSum int) int {

	count := 0

	for _, val := range array {
		if val == targetSum {
			count++
		}
	}

	return count

}

func binarySearch(array []int, targetNum int) int {
	left := 0
	right := len(array)

	found := false

	for left <= right {

		mid := right - (right-left)/2
		if array[mid] == targetNum {

			for i := mid; i >= 0; i-- {
				if array[i] == targetNum {
					left = i

				}
			}

			for i := mid; i < len(array); i++ {
				if array[i] == targetNum {
					right = i
				}
			}

			found = true
			break

		} else if array[mid] > targetNum {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if found {
		return right - left + 1
	}

	return 0
}

func binarySearchOptimised(array []int, targetNum int, searchLeft bool) int {

	left := 0
	right := len(array)

	result := 0

	for left <= right {
		mid := right - (right-left)/2

		if array[mid] == targetNum {
			result = mid

			switch searchLeft {
			case true:
				right = mid - 1
			case false:
				left = mid + 1
			}

		} else if array[mid] > targetNum {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return result

}
