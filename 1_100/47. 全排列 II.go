package main

import "sort"

func main() {

}

func permuteUnique(nums []int) [][]int {

	res := &[][]int{}
	sort.Ints(nums)
	dfs47(nums, []int{}, res, nums)
	return *res
}

func dfs47(nums []int, temp []int, res *[][]int, r []int) {
	if len(temp) == len(nums) {
		b := make([]int, len(temp))
		copy(b, temp)
		*res = append(*res, b)
	}

	for i := 0; i < len(r); i++ {
		if i > 0 && r[i] == r[i-1] {
			continue
		}
		temp = append(temp, r[i])
		b := make([]int, len(r))
		copy(b, r)
		b = append(b[:i], b[i+1:]...)
		dfs47(nums, temp, res, b)
		temp = append(temp[:len(temp)-1])
	}
}
