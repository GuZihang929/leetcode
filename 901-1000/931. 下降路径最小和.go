package main

import "fmt"

func main() {

	fmt.Println(minFallingPathSum([][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}))

}

func minFallingPathSum(matrix [][]int) int {

	length := len(matrix)
	res := make([][]int, length)
	for i, _ := range matrix[0] {
		res[i] = make([]int, length)
	}
	for i, i2 := range matrix[0] {
		res[0][i] = i2
	}

	for i := 1; i < length; i++ {
		res[i][length-1] = min(res[i-1][length-1], res[i-1][length-2]) + matrix[i][length-1]
		for j := length - 2; j > 0; j-- {
			res[i][j] = min(res[i-1][j+1], min(res[i-1][j], res[i-1][j-1])) + matrix[i][j]
		}

		res[i][0] = min(res[i-1][0], res[i-1][1]) + matrix[i][0]

		fmt.Println(res)
	}

	r := res[length-1][0]

	for _, re := range res[length-1] {
		r = min(re, r)
	}
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
