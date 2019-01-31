/*Invert a Binary Tree, Recursive Solution*/
package main

// TreeNode ... Definition of Binary Tree in Go //
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertedBinaryTree(root *TreeNode) *TreeNode {
	find(root)
	return root
}

func find(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	find(root.Left)
	find(root.Right)
}

func main() {

}
