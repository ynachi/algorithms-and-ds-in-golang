package binarytree

type Node struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

type BinarySearchTree struct {
	root *Node
}

// Init initialize the binary tree
func (t *Node) Init() {
	t.root = nil
}

// InitWithData initializes the data members of BinaryTree using the given data
func (t *Node) InitWithData(data int) {
	t.root = &Node{
		data:  data,
		left:  nil,
		right: nil,
	}
}

// FindData is used to find a data x in the BSD.
func (t *BinarySearchTree) FindData(x int) bool {
	if t.root.left == t.root.right
}
