package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	return bfs1161(root)
}

func bfs1161(root *TreeNode) int {
	if root == nil {
		return 0
	}

	m := math.MinInt
	index := 1
	res := 0
	temp := 0

	t := []*TreeNode{root}
	t1 := []*TreeNode{}

	for len(t) > 0 {
		for i := 0; i < len(t); i++ {
			temp += t[i].Val
			if t[i].Left != nil {
				t1 = append(t1, t[i].Left)
			}

			if t[i].Right != nil {
				t1 = append(t1, t[i].Right)
			}
		}

		if a, ok := max(temp, m); ok {
			res = index
			m = a
		}
		temp = 0
		index++

		t, t1 = t1, []*TreeNode{}
	}

	return res

}

func max(a, b int) (int, bool) {
	if a > b {
		return a, true
	}
	return b, false
}
