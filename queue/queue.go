package queue

import (
	"errors"
	"fmt"
)

// Queue type struct
type Queue[T any] struct {
	data []T
}

// InitQueue creates a queue and initialize it's initial capacity
func InitQueue[T any](initCap int) Queue[T] {
	return Queue[T]{
		data: make([]T, 0, initCap),
	}
}

// Size() gets the size of a queue
func (q *Queue[T]) Size() int {
	return len(q.data)
}

// IsEmpty() checks if a Queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Enqueue(value T)  Push an element into the Queue
func (q *Queue[T]) Enqueue(value T) {
	q.data = append(q.data, value)
}

// Dequeue (Remove an element from the Queue)
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var noop T // I want to return the zero value of type T
		return noop, errors.New("queue is empty")
	}
	result := q.data[0]
	q.data = q.data[1:] // Beware this is not working. I let it here for ref.
	return result, nil

}

// Print() function will print the elements of the stack
func (q *Queue[T]) String() string {
	out := ""
	if q.IsEmpty() {
		return out
	}
	for i, v := range q.data {
		if i == 0 {
			out += fmt.Sprintf("%v", v)
		} else {
			out += fmt.Sprintf(" %v", v)
		}
	}
	return out
}
