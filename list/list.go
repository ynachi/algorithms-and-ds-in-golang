/**
Package list represents a linked list with basic functions.

Add : add an element to the list
Delete : delete an element
Search : searching an element in a list
RotateRight y : rotate a list by y elements. See LeetCode 61.
**/
package list

import "strconv"

type List struct {
	val  int
	next *List
}

// Search value x in List l
// Return the first pointer holding value x or nil if not found
func (l *List) Search(x int) *List {
	if l == nil {
		return nil
	}
	if l.val == x {
		return l
	} else {
		return l.next.Search(x)
	}
}

// Insert a new node in the beginning of the list
func (l *List) Insert(x int) {
	newNode := &List{x, nil}
	newNode.next = l
	l = newNode
}

// PushBack insert a new node at the back  of the list
func (l *List) PushBack(x int) {
	p := List{x, nil}
	if l == nil {
		l = &p
	} else {
		currentNode := l
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = &p
	}
}

// Delete an item from a list
func (l *List) Delete(x int) {
	if l == nil || (l.val == x && l.next == nil) {
		l = nil
		return
	}
	tmp := l
	for tmp.next != nil && tmp.next.val != x {
		tmp = tmp.next
	}
	if tmp.next != nil {
		newNext := tmp.next
		tmp.next = tmp.next.next
		newNext.next = nil // break link for the pointer to not escape garbage collector
		newNext = nil      // set "deleted" pointer to nil
	}
}

// get the length of a linked list
func (l *List) Size() int {
	length := 0
	currentNode := l
	for currentNode != nil {
		length++
		currentNode = currentNode.next
	}
	return length
}

// Form a golang string from a linked list
// This method has no utility in practice, it will be used for testing and visual verifications only
func (l *List) ToString() string {
	result := ""
	currentList := l
	for currentList != nil {
		if len(result) > 0 {
			result = result + "-->" + strconv.Itoa(currentList.val)
			currentList = currentList.next
		} else {
			result = strconv.Itoa(currentList.val)
			currentList = currentList.next
		}
	}
	return result
}

// Reverse a linked list
func (l *List) Reverse() *List {
	current := l
	// Here is how to declare a nil struct
	var previous *List = nil
	var next *List
	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}
	// previous is the new head
	return previous
}

// Dectect a loop in a Linked List
func (l *List) DectectLoop() bool {
	current := l
	traversed := make(map[*List]bool)
	for current != nil {
		_, present := traversed[current]
		if present {
			return true
		}
		traversed[current] = true
		current = current.next
	}
	return false
}

// Rotate a list k digits from the right (see LC 61)
func (head *List) RotateRight(k int) *List {
	if head == nil || head.next == nil {
		return head
	}
	oldHead := head
	current := head
	length := 0
	for current.next != nil {
		length += 1
		current = current.next
	}
	length += 1
	current.next = oldHead
	steps := length - k
	if steps < 0 {
		steps = length - k%length
	}
	for i := 0; i < steps; i++ {
		current = current.next
	}
	newHead := current.next
	current.next = nil
	return newHead
}
