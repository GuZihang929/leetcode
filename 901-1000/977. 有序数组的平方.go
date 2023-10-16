package main

import "sort"

func main() {

}

func sortedSquares(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

//双指针法
/*
找到正负数分界值
一个指针向左走
一个指针向右走
*/
