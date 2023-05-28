package tree

import "fmt"

func (root *BinarySearchTree) Insert(parent **BinarySearchTree, value int) {
	if root == nil {
		*parent = &BinarySearchTree{value, nil, nil}
	} else {
		if value < root.value {
			root.left.Insert(&root.left, value)
		} else { // value >= root.value
			root.right.Insert(&root.right, value)
		}
	}
}

func (root *BinarySearchTree) Search(value int) bool {
	if root == nil {
		return false
	} else {
		if value < root.value {
			return root.left.Search(value)
		} else if value > root.value {
			return root.right.Search(value)
		} else {
			return true // value == root.value
		}
	}
}

func (root *BinarySearchTree) Min() int {
	if root == nil {
		return -1
	} else {
		if root.left == nil {
			return root.value
		} else {
			return root.left.Min()
		}
	}
}

func (root *BinarySearchTree) Max() int {
	if root == nil {
		return -1
	} else {
		if root.right == nil {
			return root.value
		} else {
			return root.right.Max()
		}
	}
}

func (root *BinarySearchTree) Remove(value int) *BinarySearchTree {
	if !root.Search(value) {
		fmt.Print("Value ", value, " not found.\n\n")
		return root
	}
	if value < root.value {
		root.left = root.left.Remove(value)
	} else if value > root.value {
		root.right = root.right.Remove(value)
	} else { // BinarySearchTree found
		if root.left == nil && root.right == nil { // primary case: node without children
			return nil
		} else if root.left != nil && root.right == nil { //second case - left: node with 1 child
			return root.left
		} else if root.left == nil && root.right != nil { //second case - right: node with 1 child
			return root.right
		} else { // third case: node with 2 children
			maxNodeInLeft := root.left.Max()
			root.value = maxNodeInLeft
			root.left = root.left.Remove(maxNodeInLeft)
			return root
		}
	}
	return root
}

func (root *BinarySearchTree) Height() int {
	var left, right int
	if root == nil {
		return -1
	} else {
		left = root.left.Height()
		right = root.right.Height()
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
}

func (root *BinarySearchTree) IsBST() bool {
	if root == nil {
		return false
	}

	if root.left != nil {
		if root.value > root.left.value {
			return root.left.IsBST()
		} else {
			return false
		}
	}

	if root.right != nil {
		if root.value < root.right.value {
			return root.right.IsBST()
		} else {
			return false
		}
	}

	return true
}

func (root *BinarySearchTree) Length() int {
	if root == nil {
		return 0
	} else {
		return 1 + root.left.Length() + root.right.Length()
	}
}

func (root *BinarySearchTree) Print(space int) {
	if root == nil {
		return
	}

	const increment = 3

	space += increment
	root.right.Print(space)
	fmt.Printf("%*s%d\n", space, "", root.value)
	root.left.Print(space)
}
