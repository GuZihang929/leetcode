package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	//res = append(res, []int{})
	dfs78(nums, []int{}, &res, 0)
	return res
}

func dfs78(nums []int, temp []int, res *[][]int, index int) {

	if len(nums)+1 == index {
		return
	}
	if true {
		b := make([]int, len(temp))
		copy(b, temp)
		*res = append(*res, b)
	}

	for i := index; i < len(nums); i++ {
		temp = append(temp, nums[i])
		dfs78(nums, temp, res, i+1)
		temp = temp[:len(temp)-1]
	}
}
