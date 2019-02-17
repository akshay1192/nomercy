package helperFunc

import "fmt"

var memoizationHash = make(map[int]int)

func init() {
	memoizationHash[1] = 0
	memoizationHash[2] = 1
}

// CalculateFibonacciHash and CalculateFibonacciHash2 yield same performance
func CalculateFibonacciHash2(n int) int {

	fmt.Println("Calling for n : ", n, " having hash : ", memoizationHash)

	if val, ok := memoizationHash[n]; ok {
		return val
	}

	if _, ok := memoizationHash[n-1]; ok {
		memoizationHash[n] = memoizationHash[n-1] + memoizationHash[n-2]
		return memoizationHash[n]
	}

	if _, ok := memoizationHash[n-2]; ok {
		memoizationHash[n] = CalculateFibonacciHash(n-1) + memoizationHash[n-2]
		return memoizationHash[n]
	}

	return CalculateFibonacciHash(n-1) + CalculateFibonacciHash(n-2)
}

func CalculateFibonacciHash(n int) int {

	if val, ok := memoizationHash[n]; ok {
		return val
	}

	memoizationHash[n] = CalculateFibonacciHash(n-1) + CalculateFibonacciHash(n-2)
	return memoizationHash[n]
}

func CalculateFibonacciWithoutHash(n int) int {

	switch n {
	case 1:
		return 0
	case 2:
		return 1
	default:
		return CalculateFibonacciWithoutHash(n-1) + CalculateFibonacciWithoutHash(n-2)
	}

}

func CalculateFibonacciLoop(n int) int {

	first := 0
	second := 1
	i := 2
	lastCalculated := 0

	for i < n {
		lastCalculated = first + second
		first = second
		second = lastCalculated
		i++
	}

	switch n {

	case 1:
		return first
	case 2:
		return second
	default:
		return lastCalculated

	}

}
