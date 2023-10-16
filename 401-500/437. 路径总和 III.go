package main

func main() {

}

func pathSum(root *TreeNode, targetSum int) int {
	a = targetSum
	return dfs437(root)
}

var a int

func dfs437(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return path(root, a) + dfs437(root.Left) + dfs437(root.Right)
}

func path(root *TreeNode, sub int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == sub {
		res++
	}

	return path(root.Left, sub-root.Val) + path(root.Right, sub-root.Val) + res
}
