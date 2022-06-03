package queue

import "fmt"

// Node type struct
type Node struct {
	value interface{}
	next  *Node
}

// Queue type struct
type Queue struct {
	tail *Node
	head *Node
	size int
}

// Get the size of a Queue
func (q *Queue) Size() int {
	return q.size
}

// Check if a Queue is empty
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Get the top element of the Queue
func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return nil
	}
	return q.head.value
}

// Enqueue (Push an element into the Queue)
func (q *Queue) Enqueue(value interface{}) {
	tmp := &Node{value, nil}
	if q.IsEmpty() {
		q.tail = tmp
		q.head = tmp
	}
	q.tail.next = tmp
	q.tail = tmp
	q.size++
}

// Dequeue (Remove an element from the Queue)
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return nil
	}
	result := q.head.value
	q.head = q.head.next
	q.size--
	return result

}

// Print() function will print the elements of the stack
func (q *Queue) Print() {
	currentNode := q.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Println()
}
