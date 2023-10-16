package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(longestZigZag(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	return dfs1372(root, "", 0) - 1
}

// cum为前面累计的数量
func dfs1372(root *TreeNode, direction string, cum int) int {
	if root == nil {
		return cum
	}
	var l, r int
	if direction == "right" {
		l = dfs1372(root.Left, "left", cum+1)
		// 当前位置重新开始计数
		r = dfs1372(root.Right, "right", 1)
	} else if direction == "left" {
		l = dfs1372(root.Left, "left", 1)
		r = dfs1372(root.Right, "right", cum+1)
	} else {
		l = dfs1372(root.Left, "left", 1)
		r = dfs1372(root.Right, "right", 1)
	}
	if l > r {
		return l
	} else {
		return r
	}
}
