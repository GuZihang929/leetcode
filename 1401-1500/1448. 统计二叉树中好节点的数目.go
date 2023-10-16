package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	res := 0
	dfs1448(root, root.Val, &res)
	return res
}

func dfs1448(root *TreeNode, max int, res *int) {
	if root == nil {
		return
	}

	if root.Val >= max {
		max = root.Val
		*res++
	}

	dfs1448(root.Left, max, res)
	dfs1448(root.Right, max, res)
}
