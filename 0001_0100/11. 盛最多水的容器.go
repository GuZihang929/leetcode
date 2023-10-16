package main

func main() {

}

func maxArea(height []int) int {
	Max := Area(min(height[0], height[len(height)-1]), len(height)-1)
	left, right := 0, len(height)-1
	for left < right {
		leng := right - left
		Max = max(Max, Area(leng, min(height[left], height[right])))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return Max

}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Area(x, y int) int {
	return x * y
}
