// Equilibrium in an array : https://www.youtube.com/watch?v=W-t1rjLxvQw
// Ex : [-7 1 5 2 -4 3 0]    3
// Ex : [-7 1 5 2 -1 3 0]	 5

package main

import "fmt"

func main() {
	intSlice := []int{-7, 1, 5, 2, -1, 3, 0}

	fmt.Println("Naive approach : ", naiveApproach(intSlice))

	fmt.Println("Better approach : ", betterAppraoch(intSlice))

	fmt.Println("Best approach : ", bestApproach(intSlice))

}

// O(n2) time and O(1) space
func naiveApproach(intSlice []int) int {
	for i := 1; i < len(intSlice); i++ {

		leftSum, rightSum := 0, 0

		for j := 0; j < i; j++ {
			leftSum += intSlice[j]
		}

		for j := i + 1; j < len(intSlice); j++ {
			rightSum += intSlice[j]
		}

		if leftSum == rightSum {
			return i
		}

	}

	return -1
}

// O(n) time and O(n) space
func betterAppraoch(intSlice []int) int {

	len := len(intSlice)
	leftSum, rightSum := make([]int, len), make([]int, len)

	leftSum[0] = 0
	rightSum[len-1] = 0

	for i := 1; i < len; i++ {
		leftSum[i] = leftSum[i-1] + intSlice[i-1]
	}

	for i := len - 2; i >= 0; i-- {
		rightSum[i] = rightSum[i+1] + intSlice[i+1]

		if rightSum[i] == leftSum[i] {
			return i
		}
	}

	return -1
}

// O(n) time and O(1) space
func bestApproach(intSlice []int) int {
	len := len(intSlice)
	sum,leftSum := 0,0

	for i := 0; i < len; i++ {
		sum += intSlice[i]
	}

	for i := 1; i < len-1; i++ {

		leftSum += intSlice[i-1]
		rightSum := sum - (leftSum + intSlice[i])

		if rightSum == leftSum {
			return i
		}
	}

	return -1

}
