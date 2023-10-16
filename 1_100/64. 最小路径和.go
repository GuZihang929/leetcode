package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	hight := len(grid)
	length := len(grid[0])
	res := make([][]int, hight)
	for i := range res {
		res[i] = make([]int, length) // 每一行4列
	}
	fmt.Println(res)
	res[0][0] = grid[0][0]
	for i := 1; i < hight; i++ {
		res[i][0] = grid[i][0] + res[i-1][0]
	}

	for i := 1; i < length; i++ {
		res[0][i] = grid[0][i] + res[0][i-1]
	}
	fmt.Println(res)
	for i := 1; i < hight; i++ {
		for j := 1; j < length; j++ {
			res[i][j] = min64(res[i-1][j], res[i][j-1]) + grid[i][j]
		}

	}
	fmt.Println(res)
	return res[hight-1][length-1]

}

func min64(a, b int) int {
	if a < b {
		return a
	}
	return b
}
