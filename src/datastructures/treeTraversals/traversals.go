package main

import "fmt"

// Node -> Binary Tree Representation
type Node struct {
	val   int
	left  *Node
	right *Node
}

// InOrder -> Traverses tree in order
func InOrder(root *Node) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Printf("%d ", root.val)
	InOrder(root.right)
}

// PreOrder -> Traverses tree pre order
func PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.val)
	PreOrder(root.left)
	PreOrder(root.right)
}

// PostOrder -> Traverses tree in post order
func PostOrder(root *Node) {
	if root == nil {
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Printf("%d ", root.val)
}

func main() {

	// Binary Tree Representation
	root := &Node{10, nil, nil}
	root.left = &Node{20, nil, nil}
	root.right = &Node{30, nil, nil}
	root.left.left = &Node{40, nil, nil}
	root.left.right = &Node{50, nil, nil}
	root.right.right = &Node{60, nil, nil}
	root.left.left.left = &Node{70, nil, nil}

	// InOrder Traversal
	fmt.Println("InOrder: ")
	InOrder(root)
	// PreOrder Traversal
	fmt.Println(" ")
	fmt.Println("PreOrder: ")
	PreOrder(root)
	// PostOrder Traversal
	fmt.Println(" ")
	fmt.Println("PostOrder: ")
	PostOrder(root)
	fmt.Println(" ")

}
