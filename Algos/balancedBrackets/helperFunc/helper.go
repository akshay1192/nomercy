package helperFunc

func CheckBalancedBrackets(s string) bool {

	var stack []rune

	mapping := map[rune]rune {
		'}' : '{',
		')' : '(',
		']' : '[',
	}

	for _,char := range s {

		if char == 123 || char == 91 || char == 40 {
			stack = append(stack,char)
		}

		if char == '}' || char == ']' || char == ')' {

			if len(stack) == 0 {
				return false
			}

			if mapping[char] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}

	}

	return len(stack) == 0

}
