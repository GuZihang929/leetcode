package main

import "fmt"

func main() {

	fmt.Println(rob([]int{0, 1, 0, 3}))
}

func rob(nums []int) int {
	length := len(nums)

	if length == 1 {
		return nums[0]
	}

	if length == 2 {
		return Max(nums[0], nums[1])
	}
	res := make([]int, length)
	res[0] = nums[0]
	res[1] = Max(nums[1], res[0])

	for i := 2; i < length; i++ {
		res[i] = Max(res[i-1], res[i-2]+nums[i])

	}
	fmt.Println(res)
	return res[length-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
