package main

import "fmt"

func main() {
	tree := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	fmt.Println(diameterOfBinaryTree(&tree))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {

	result := 0
	MaxTree(root, &result)
	return result
}

func MaxTree(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := MaxTree(root.Left, result)
	right := MaxTree(root.Right, result)
	*result = Max(*result, left+right)
	return Max(left, right) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
