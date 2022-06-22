package binarytree

type Node struct {
	data   int
	parent *Node
	left   *Node
	right  *Node
}

type BinaryTree struct {
	root *Node
}

// FindData is used to find a data x in the BSD.
func (t *BinaryTree) FindData(x int) (*Node, bool) {
	if t.root == nil {
		return nil, false
	}
	if t.root.data == x {
		return t.root, true
	}
	if x < t.root.data {
		temp := BinaryTree{t.root.left}
		return temp.FindData(x)
	} else {
		temp := BinaryTree{t.root.right}
		return temp.FindData(x)
	}
}

// InsertData is used to insert data x in the BSD.
func (t *BinaryTree) InsertData(x int) {
	tempNode := &Node{
		data: x,
	}
	if t.root == nil {
		t.root = tempNode
		return
	}
	currentNode := t.root
	var direction string // which one from parent.left or parent.right is will be null ?
	var currentParent *Node
	for currentNode != nil {
		currentParent = currentNode.parent
		if x < currentNode.data {
			currentNode = currentNode.left
			direction = "left"
		} else {
			currentNode = currentNode.right
			direction = "right"
		}
	}
	if direction == "left" {
		currentParent.left = tempNode
	} else {
		currentParent.right = tempNode
	}
}
