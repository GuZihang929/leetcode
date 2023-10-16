package main

import (
	"strconv"
	"strings"
)

func main() {

}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	s := zhong(root)
	s2 := zhong(subRoot)
	return strings.Contains(s, s2)
}

func zhong(r *TreeNode) string {
	if r == nil {
		return "."
	}
	s := zhong(r.Left)
	itoa := strconv.Itoa(r.Val)
	s1 := zhong(r.Right)

	return s + "L" + itoa + s1 + "R"
}
