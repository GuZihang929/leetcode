package main

import "fmt"

func main() {

	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}

func minimumTotal(triangle [][]int) int {

	length := len(triangle)

	res := make([][]int, length)

	for i, _ := range res {
		res[i] = make([]int, i+1)
	}
	res[0][0] = triangle[0][0]
	for i := 1; i < length; i++ {
		res[i][0] = res[i-1][0] + triangle[i][0]
		res[i][i] = res[i-1][i-1] + triangle[i][i]
	}

	for i := 2; i < length; i++ {
		for j := 1; j < i; j++ {
			res[i][j] = min(res[i-1][j], res[i-1][j-1]) + triangle[i][j]
		}
	}
	re := res[length-1][0]

	for i := 1; i < length; i++ {
		re = min(re, res[length-1][i])
	}

	fmt.Println(res)

	return re
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
