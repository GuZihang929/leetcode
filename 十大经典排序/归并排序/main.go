package main

import "fmt"

func main() {
	fmt.Println(MergeSort([]int{4, 3, 12, 5, 124, 465, 123, 53, 675, 12325, 15, 61, 56, 234, 2}))

}

func MergeSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	left := MergeSort(arr[:len(arr)/2])
	right := MergeSort(arr[len(arr)/2:])
	return Merge(left, right)

}

func Merge(left, right []int) []int {
	res := make([]int, len(left)+len(right))
	i, j := 0, 0
	m := 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res[m] = left[i]
			i++
		} else {
			res[m] = right[j]
			j++
		}
		m++
	}
	if i == len(left) {
		res = append(res[0:m], right[j:]...)
	}
	if j == len(right) {
		res = append(res[0:m], left[i:]...)
	}
	return res
}
