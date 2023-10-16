package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {

	res := &[][]int{}
	dfs46(nums, []int{}, res, nums)
	return *res
}

func dfs46(nums []int, temp []int, res *[][]int, r []int) {
	if len(temp) == len(nums) {
		b := make([]int, len(temp))
		copy(b, temp)
		*res = append(*res, b)
	}

	for i := 0; i < len(r); i++ {
		temp = append(temp, r[i])
		b := make([]int, len(r))
		copy(b, r)
		b = append(b[:i], b[i+1:]...)
		dfs46(nums, temp, res, b)
		temp = append(temp[:len(temp)-1])
	}
}
