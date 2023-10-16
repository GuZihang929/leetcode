package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(5, 29))
}

func combinationSum3(k int, n int) [][]int {
	temp, res := []int{}, &[][]int{}

	dsf216(temp, res, k, n, 1)
	return *res
}

func dsf216(temp []int, res *[][]int, k, n int, index int) {
	if k < 0 || n < 0 {
		return
	}
	if k == 0 && n == 0 {
		b := make([]int, len(temp))
		copy(b, temp)
		*res = append(*res, b)
		return
	}

	for i := index; i <= 9; i++ {
		temp = append(temp, i)
		dsf216(temp, res, k-1, n-i, i+1)
		temp = temp[:len(temp)-1]

	}
}
