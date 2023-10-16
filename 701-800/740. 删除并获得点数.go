package main

import "fmt"

func main() {
	deleteAndEarn([]int{3, 1})
}

func deleteAndEarn(nums []int) int {

	maxVal := 0
	for _, val := range nums {
		maxVal = Max(maxVal, val)
	}
	sum := make([]int, maxVal+1)
	for _, val := range nums {
		sum[val] += val
	}
	fmt.Println(sum)
	return rob(sum)

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
	res[1] = Max(res[0], nums[1])

	for i := 2; i < length; i++ {
		res[i] = Max(res[i-1], res[i-2]+nums[i])

	}
	return res[length-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
