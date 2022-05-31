package queue

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
	return nil
}

// Enqueue (Push an element into the Queue)
func (q *Queue) Enqueue(value interface{}) {

}

// Dequeue (Remove an element from the Queue)
func (q *Queue) Dequeue() interface{} {
	return nil
}

// Print() function will print the elements of the stack
func (s *Queue) Print() {
}
