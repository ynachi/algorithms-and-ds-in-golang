/**
Package list represents a linked list with basic functions.

Add : add an element to the list
Delete : delete an element
Search : searching an element in a list
RotateRight y : rotate a list by y elements. See LeetCode 61.
**/
package list

type List struct {
	Val  int
	Next *List
}

// Search value x in List l
// Return the first pointer holding value x or nil if not found
func (l *List) Search(x int) *List {
	if l == nil {
		return nil
	}
	if l.Val == x {
		return l
	} else {
		return l.Next.Search(x)
	}
}

// Insert a new node in the begining of the list
func (l *List) Insert(x int) {
	newNode := &List{x, nil}
	newNode.Next = l
	l = newNode
}

// Insert a new node at the back  of the list
func (l *List) PushBack(x int) {
	p := List{x, nil}
	if l == nil {
		l = &p
	} else {
		currentNode := l
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = &p
	}
}

// Delete an item from a list
func (l *List) Delete(x int) {
	if l == nil || (l.Val == x && l.Next == nil) {
		l = nil
		return
	}
	tmp := l
	for tmp.Next != nil && tmp.Next.Val != x {
		tmp = tmp.Next
	}
	if tmp.Next != nil {
		newNext := tmp.Next
		tmp.Next = tmp.Next.Next
		newNext.Next = nil // break link for the pointer to not escape garbage collector
		newNext = nil      // set "deleted" pointer to nil
	}
}

// Form a golang slice from a linked list
// This method has no utility in practice, it will be used for testing and visual verifications only
func (l *List) ToSlice() []int {
	var result []int
	currentList := l
	for currentList != nil {
		result = append(result, currentList.Val)
		currentList = currentList.Next
	}
	return result
}

// Rotate a list k digits from the right (see LC 61)
func (*List) RotateRight(k int) *List {
	return nil
}
