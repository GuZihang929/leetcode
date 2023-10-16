package main

import "fmt"

func main() {

	fmt.Println(productExceptSelf01([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))
	h := make([]int, len(nums))

	pre[0] = 1
	h[len(nums)-1] = 1
	for i := 0; i < len(nums)-1; i++ {
		pre[i+1] = nums[i] * pre[i]
		h[len(nums)-2-i] = nums[len(nums)-1-i] * h[len(nums)-1-i]
	}
	//fmt.Println(pre)
	//fmt.Println(h)

	res := make([]int, len(nums))

	for i := 0; i < len(res); i++ {
		res[i] = pre[i] * h[i]
	}
	return res
}

func productExceptSelf01(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	fmt.Println(res)
	temp := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		res[i] *= temp
		temp *= nums[i]

	}
	return res
}
