package main

func main() {

}

func longestSubarray(nums []int) int {
	k := 1
	ans := 0
	var left, lsum, rsum int
	for right, v := range nums {
		rsum += 1 - v
		for lsum < rsum-k {
			lsum += 1 - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
