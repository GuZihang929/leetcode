package main

func main() {

}

func findLengthOfLCIS(nums []int) int {
	max := 1
	temp := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			temp++
		} else {
			temp = 0
		}
		if max < temp {
			max = temp
		}

	}
	return max
}
