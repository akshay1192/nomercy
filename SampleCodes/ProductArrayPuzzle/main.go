//Product array puzzle :
//https://www.youtube.com/watch?v=R745JLPox_A
//https://www.geeksforgeeks.org/a-product-array-puzzle/

// Condition : do not use division operator
package main

import (
	"fmt"
)

func main() {

	intSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("Using division operator : ", usingDivisionOperator(intSlice))
	fmt.Println("Using naive approach : ", naiveApproach(intSlice))
	fmt.Println("Using better approach : ", betterApproach(intSlice))
	fmt.Println("Using best approach : ", bestApproach(intSlice))
}

// O(n) time and O(n) space
func usingDivisionOperator(intSlice []int) []int {
	prodArray := make([]int, len(intSlice), 5)
	prod := 1

	for _, val := range intSlice {
		prod *= val
	}

	for i := 0; i < len(intSlice); i++ {
		prodArray[i] = prod / intSlice[i]
	}

	return prodArray
}

// O(n2) time and O(n) space
func naiveApproach(intSlice []int) []int {

	len := len(intSlice)
	var prodArray []int

	for i := 0; i < len; i++ {

		prod := 1
		for j := 0; j < len; j++ {

			if i == j {
				continue
			}
			prod *= intSlice[j]
		}

		prodArray = append(prodArray, prod)
	}

	return prodArray
}

// O(n) time, O(n) space and O(n) auxiliary space
// Aux space is for storing left and right prod
func betterApproach(intSlice []int) []int {

	leftProd, rightProd, prodArray := make([]int, 5, 5), make([]int, 5, 5), make([]int, 5, 5)

	leftProd[0] = 1
	rightProd[len(intSlice)-1] = 1

	// left prod array till i excluding element at index i
	for i := 1; i < len(intSlice); i++ {
		leftProd[i] = leftProd[i-1] * intSlice[i-1]
	}

	// right prod array from i excluding element at index i
	for i := len(intSlice) - 2; i >= 0; i-- {
		rightProd[i] = rightProd[i+1] * intSlice[i+1]
	}

	// left[i] * right[i] = prod[i] at i excluding element i
	for i := 0; i < len(intSlice); i++ {
		prodArray[i] = leftProd[i] * rightProd[i]
	}

	return prodArray
}

// O(n) time, O(n) space without any auxiliary space
// prod array is saving the result
func bestApproach(intSlice []int) []int {

	prodArray := make([]int, 5, 5)
	temp := 1

	// left prod array till i excluding element at index i
	// populating the left prod in prodArray
	for i := 0; i < len(intSlice); i++ {

		// note that there is not *= here, just =
		prodArray[i] = temp
		temp *= intSlice[i]
	}

	temp = 1

	// prodArray already had left prod and now we are multiply right prod at i
	for i := len(intSlice) - 1; i >= 0; i-- {

		// note that there is *= here, not just = as we already have prodArray[i] populated in upper loop
		prodArray[i] *= temp
		temp *= intSlice[i]
	}

	return prodArray

}
