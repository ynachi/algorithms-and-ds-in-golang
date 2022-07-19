package stack

import (
	"errors"
	"fmt"
)

// Stack type struct
type Stack[T any] struct {
	value []T
}

// InitStack creates a stack and initialize it's initial capacity
func InitStack[T any](initCap int) Stack[T] {
	return Stack[T]{
		value: make([]T, 0, initCap),
	}
}

// Get the size of the Stack
func (s *Stack[T]) Size() int {
	return len(s.value)
}

// Check if the Stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.value) == 0
}

// Peek() returns the top value of the stack
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var noop T // I want to return the zero value of type T
		return noop, errors.New("stack is empty")
	}
	return s.value[len(s.value)-1], nil
}

// Push() function  will push the value into the stack
func (s *Stack[T]) Push(data T) {
	s.value = append(s.value, data)
}

// Pop() function will pop the value from the top of the stack
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var noop T // I want to return the zero value of type T
		return noop, errors.New("stack is empty")
	}
	result := s.value[len(s.value)-1]
	s.value = s.value[:len(s.value)-1]
	return result, nil
}

// Print() function will print the elements of the stack
func (s *Stack[T]) String() string {
	out := ""
	if s.IsEmpty() {
		return out
	}
	for i, v := range s.value {
		if i == 0 {
			out += fmt.Sprintf("%v", v)
		} else {
			out += fmt.Sprintf(" %v", v)
		}
	}
	return out
}

/**
  Now let's practice stacks
**/

// BalancedParentheses(str) check if a string of parentheses is balanced
// Which means every opening parenthese has it's equivalent closing one
