package main

import "fmt"

func main() {
	fmt.Println(findCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}))
}

func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	visited := make([]bool, len(M))
	res := 0
	for i := range M {
		if !visited[i] {
			dfs547(M, i, visited)
			res++
		}
	}
	return res
}

func dfs547(M [][]int, cur int, visited []bool) {
	visited[cur] = true
	for j := 0; j < len(M[cur]); j++ {
		if !visited[j] && M[cur][j] == 1 {
			dfs547(M, j, visited)
		}
	}
}
