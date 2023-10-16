package main

import "fmt"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 8}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	temp := make([]int, m+n)

	copy(temp, nums1)

	temp = temp[:m]

	for i, j := 0, 0; i < m && j < n; {
		if temp[0] < nums2[0] {
			nums1[i+j] = temp[0]
			temp = temp[1:]
			i++
		} else {
			nums1[i+j] = nums2[0]
			nums2 = nums2[1:]
			j++
		}
	}
	if len(nums2) > 0 {
		temp = nums2
	}
	j := 0
	for i := len(temp) - 1; i > -1; i-- {
		nums1[m+n-1-j] = temp[i]
		j++
	}
	fmt.Println(nums1)
}
