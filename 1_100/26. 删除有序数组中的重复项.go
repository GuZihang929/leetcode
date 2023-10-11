package main

func main() {

}

func removeDuplicatesI(nums []int) int {

	temp := make(map[int]int)

	for i := 0; i < len(nums); {
		temp[nums[i]]++
		if temp[nums[i]] > 1 {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}
