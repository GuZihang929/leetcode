package main

import "fmt"

func main() {
	fmt.Println(combine(5, 2))

}

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	dfs(&res, c, k, n, 1)
	return res
}

func dfs(res *[][]int, cand []int, k, n, m int) {
	if len(cand) == k {
		b := make([]int, len(cand))
		copy(b, cand)
		*res = append(*res, b)
		//*res = append(*res, cand)
		return
	}
	for i := m; i <= n-(k-len(cand))+1; i++ {
		cand = append(cand, i)
		dfs(res, cand, k, n, i+1)
		cand = cand[:len(cand)-1]
	}
}
