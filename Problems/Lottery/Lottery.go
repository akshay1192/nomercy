// You will be given 2 number a and b and you need to calculate total number of ways to maximise the win chances

// Example 1 : a = 3 and b = 12,   3 4 5 6 7 8 9 10 11 12 13 = 3 4 5 6 7 8 9 1 2 3 4 (10 => 1 + 0 = 1 )
// 			   3,4 occurring twice and rest numbers are occurring only once so answer would be
// 			   ( Total ways : 2 (Lottery number = 3 and 4) , distributing amongst 2 winners as max occurrences have value = 2)
// Output : 2 ,2

// Example 2 : a = 1 and b = 5,   1 2 3 4 5
//			   all numbers are occurring only one  (Total ways : 5 (All the numbers), distributing amongst 1 winner
//Output : 5,1

package main

import (
	"fmt"
	"sort"
)

func calculateSum(num int) int {

	sum := 0

	for num != 0 {
		sum += num % 10
		num = num / 10
	}

	return sum
}

func main() {

	var a, b int
	fmt.Scanln(&a, &b)

	occurenceMap := map[int]int{}

	for i := a; i <= b; i++ {

		/* using util func
		sum := calculateSum(i)*/

		// using anonymous functions
		sum := func(num int) int {

			x := 0

			for num != 0 {
				x += num % 10
				num = num / 10
			}
			return x
		}(i)

		if val, ok := occurenceMap[sum]; ok {
			val++
			occurenceMap[sum] = val
		} else {
			occurenceMap[sum] = 1
		}

	}

	occurrenceCount := []int{}
	for _, v := range occurenceMap {
		occurrenceCount = append(occurrenceCount, v)
	}

	// sorting with compare func
	sort.Slice(occurrenceCount, func(i, j int) bool {
		return occurrenceCount[i] > occurrenceCount[j]
	})

	//fmt.Println(occurrenceCount)

	numberOfWays := 0

	for i, _ := range occurrenceCount {
		if occurrenceCount[i] == occurrenceCount[0] {
			numberOfWays++
		} else {
			break
		}
	}

	fmt.Println("Total number of ways to win lottery : ", numberOfWays, " having ", occurrenceCount[0], " winner(s) each")

}
