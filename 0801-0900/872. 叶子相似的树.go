package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	temp := dfs872(root1, &[]int{})
	temp1 := dfs872(root2, &[]int{})
	if len(*temp) != len(*temp1) {
		return false
	}

	for i := 0; i < len(*temp); i++ {
		if (*temp)[i] != (*temp1)[i] {
			return false
		}
	}

	return true
}

func dfs872(root *TreeNode, res *[]int) *[]int {
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		*res = append(*res, root.Val)
		return res
	}

	dfs872(root.Left, res)
	dfs872(root.Right, res)
	return res
}

func isDfs872(root *TreeNode, res *[]int) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		temp := (*res)[0]
		*res = (*res)[1:]
		return root.Val == temp
	}

	return isDfs872(root.Left, res) && isDfs872(root.Right, res)

}
