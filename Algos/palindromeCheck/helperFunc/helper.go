package helperFunc

import "golang.org/x/exp/errors/fmt"

func IsPalindrome(input string) {

	if input == "" {
		return
	}

	fmt.Println("Palindrome reverse string : ", palindromeReverse(input))
	fmt.Println("Palindrome reverse without string creation : ", palindromeReverseCheck(input))
	fmt.Println("Palindrome recursion : ", palindromeRecursion(input, 0))
	fmt.Println("Palindrome half : ", palindromeHalf(input))

}

// O(N^2) time and O(N) space
func palindromeReverse(input string) bool {
	reverseString := []byte{}

	for i := len(input) - 1; i >= 0; i-- {
		reverseString = append(reverseString, input[i])
	}

	return string(reverseString) == input
}

// O(N) time and O(N) space
func palindromeReverseCheck(input string) bool {
	dupString := input
	isPalindrome := true
	length := len(input)

	for i := 0; i < length; i++ {
		if input[i] != dupString[length-i-1] {
			isPalindrome = false
			break
		}
	}

	return isPalindrome
}

// O(N) time and O(N) space
func palindromeRecursion(input string, currentIndex int) bool {
	if currentIndex == len(input)-1 {
		return true
	}

	checked := palindromeRecursion(input, currentIndex+1)
	return checked && input[len(input)-1-currentIndex] == input[currentIndex]
}

// O(N) time and O(1) space
func palindromeHalf(input string) bool {

	len := len(input)
	isPalindrome := true

	for i := 0; i < len/2; i++ {
		if input[i] != input[len-i-1] {
			isPalindrome = false
			break
		}
	}

	return isPalindrome
}
