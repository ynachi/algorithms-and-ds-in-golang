package binarytree

type Node struct {
	data  int
	left  *Node
	right *Node
}

// FindData is used to find a data x in the BSD.
func (root *Node) FindData(x int) (*Node, bool) {
	switch {
	case root == nil:
		return nil, false
	case x < root.data:
		return root.left.FindData(x)
	case x > root.data:
		return root.right.FindData(x)
	default:
		return root, true
	}
}

// InsertData is used to insert data x in the BSD.
func (root *Node) InsertData(x int) *Node {
	tempNode := &Node{
		data: x,
	}
	if root == nil {
		return tempNode
	}
	if x < root.data {
		root.left = root.left.InsertData(x)
	} else if x > root.data {
		root.right = root.right.InsertData(x)
	}
	return root
}

// IsBST checks if a tree is a binary search tree
func (root *Node) IsBST() bool {
	return isBSTHelper(root, nil)
}

// Leveraging inorder traversal
func isBSTHelper(root *Node, prev *Node) bool {
	// Base case
	if root == nil {
		return true
	}
	if prev != nil && prev.data >= root.data {
		return false
	}
	// Go left, go node, go right
	//	return isBSTHelper(root.left, root) && isBSTHelper(root, nil) && isBSTHelper(root.right, root)
	return isBSTHelper(root.left, root) && isBSTHelper(root.right, root)
}
