package main

func main() {

}

func removeDuplicatesII(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow = fast
			fast++
		} else {
			if fast-slow == 2 {
				nums = append(nums[:slow], nums[slow+1:]...)
			} else {
				fast++
			}
		}
	}
	return len(nums)
}
