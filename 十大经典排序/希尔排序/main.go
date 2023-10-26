package main

import "fmt"

func main() {
	fmt.Println(ShellSort([]int{4, 3, 12, 5, 546, 123, 53, 15, 61, 56, 7, 2}))

}

func ShellSort(arr []int) []int {
	gap := len(arr) >> 1
	for gap > 0 {
		for i := 0; i < len(arr)-gap; i += gap {
			j := i
			for j >= 0 && arr[j+gap] < arr[j] {
				arr[j+gap], arr[j] = arr[j], arr[j+gap]
				j -= gap
			}
		}
		gap = gap >> 1
	}

	return arr
}
