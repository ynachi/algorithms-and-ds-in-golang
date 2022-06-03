package stack

import "fmt"

// Node type struct
type Node struct {
	value interface{}
	next  *Node
}

// Queue type struct
type Stack struct {
	head *Node
	size int
}

// Get the size of the Stack
func (s *Stack) Size() int {
	return s.size
}

// Check if the Stack is empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Peek() returns the top value of the stack
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return nil
	}
	return s.head.value
}

// Push() function  will push the value into the stack
func (s *Stack) Push(value interface{}) {
	s.head = &Node{
		value: value,
		next:  s.head,
	}
	s.size++
}

// Pop() function will pop the value from the top of the stack
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return nil
	}
	result := s.head.value
	s.head = s.head.next
	s.size--
	return result
}

// Print() function will print the elements of the stack
func (s *Stack) Print() {
	currentNode := s.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Println()
}