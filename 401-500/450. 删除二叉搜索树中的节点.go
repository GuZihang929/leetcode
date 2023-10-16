package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	r := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	a := r
	r = nil
	fmt.Println(r)
	fmt.Println(a)
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	res := root
	if root == nil {
		return nil
	}
	for root != nil && root.Val != key {
		if key > root.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	fmt.Println(root.Val)

	if root.Left == nil && root.Right != nil {
		root = root.Right
	} else if root.Left != nil && root.Right == nil {
		root = root.Left
	} else if root.Left == nil && root.Right == nil {
		root = nil
	} else {

		dfs450(root.Right, root.Left)

		root.Val = root.Right.Val
		root.Left = root.Right.Left
		root.Right = root.Right.Right
	}

	return res

}

func dfs450(root, rL *TreeNode) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		root.Left = rL
		return
	}
	dfs450(root.Left, rL)
	dfs450(root.Right, rL)
}
