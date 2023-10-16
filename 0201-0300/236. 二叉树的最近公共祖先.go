package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	nodep, _ := dfs236(root, p, []*TreeNode{})
	nodeq, _ := dfs236(root, q, []*TreeNode{})
	m := min(len(nodep), len(nodeq))

	for i := 0; i < m; i++ {
		if nodep[i] != nodeq[i] {
			return nodeq[i-1]
		}
		if i == m-1 {
			return nodep[i]
		}
	}

	return nil

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func dfs236(root, p *TreeNode, res []*TreeNode) ([]*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	res = append(res, root)
	if root == p {
		return res, true
	}
	nodes, b := dfs236(root.Left, p, res)
	if b {
		return nodes, true
	}
	treeNodes, b2 := dfs236(root.Right, p, res)
	if b2 {
		return treeNodes, true
	}
	return nil, false
}
