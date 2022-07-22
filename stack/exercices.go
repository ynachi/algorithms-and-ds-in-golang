package stack

import (
	"errors"
	"strconv"
)

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

// DecToBin(str) converts a decimal number into binary
func DecToBin(number int) string {
	binaryStk := InitStack[int](10)
	tmp := number
	for tmp > 0 {
		binaryStk.Push(tmp % 2)
		tmp = tmp / 2
	}
	result := ""
	for !binaryStk.IsEmpty() {
		digit, _ := binaryStk.Pop()
		result = result + strconv.Itoa(digit)
	}
	if result == "" {
		result = "0"
	}
	return result
}

// EvaluatePostfix evaluates a postfix expression given as a string
// and return its value. For simpliciy, / operator is not included.

// Not working for now
func EvaluatePostfix(expr string) (int, error) {
	operators := map[rune]bool{
		'+': true,
		'-': true,
		'*': true,
	}
	stk := InitStack[int](15)
	for _, digit := range expr {
		if _, ok := operators[digit]; !ok {
			// Beware int(digit) won't work as it just convert the ascii code to int
			stk.Push(int(digit - '0'))
			continue
		}
		var intermediateComputation int
		d1, err := stk.Pop()
		if err != nil {
			return -1, err
		}
		d2, err := stk.Pop()
		if err != nil {
			return -1, err
		}
		switch digit {
		case '+':
			intermediateComputation = d2 + d1
		case '-':
			intermediateComputation = d2 - d1
		case '*':
			intermediateComputation = d2 * d1
		}
		stk.Push(intermediateComputation)
	}
	result, _ := stk.Pop()
	if stk.IsEmpty() {
		return result, nil
	} else {
		return -1, errors.New("invalid expression")
	}
}
