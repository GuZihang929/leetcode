package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	//temp := []int{}
	nums = append(nums, 1)
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		dfs90(nums, []int{nums[i]}, &res, i+1)
	}
	return res
}

func dfs90(nums []int, temp []int, res *[][]int, index int) {

	if len(nums) == index {
		return
	}
	if true {
		b := make([]int, len(temp))
		copy(b, temp)
		*res = append(*res, b)
	}
	for i := index; i < len(nums); i++ {
		temp = append(temp, nums[i])
		dfs90(nums, temp, res, i+1)
		temp = temp[:len(temp)-1]
	}

}
