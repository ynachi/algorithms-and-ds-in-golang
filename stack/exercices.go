package stack

/**
  Now let's practice stacks
**/

// BalancedParentheses(str) check if a string of parentheses is balanced
// Which means every opening parenthese has it's equivalent closing one
// We assume that the input string contains only '(' and ')'.
func BalancedParentheses(in string) bool {
	stk := InitStack[rune](10)
	for _, v := range in {
		if v == '(' {
			stk.Push(v)
		} else {
			_, err := stk.Pop()
			if err != nil {
				return false
			}
		}
	}
	return stk.IsEmpty()
}

// BalancedSymbols(str) is a variant of balanced parentheses
// But with different symbols (), [], {}
func BalancedSymbols(in string) bool {
	symbols := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stk := InitStack[rune](10)
	for _, char := range in {
		if _, ok := symbols[char]; ok {
			stk.Push(char)
			continue
		}
		opening, err := stk.Pop()
		// if the closing of the poped char is not the good one
		// or if there are some error (which is trying to pop an
		// an empty stack in our case
		if err != nil || symbols[opening] != char {
			return false
		}
	}
	return stk.IsEmpty()
}
