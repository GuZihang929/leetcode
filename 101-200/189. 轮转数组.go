package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	Count int
}

func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c)
}

func foo(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

// 解法一 时间复杂度 O(n)，空间复杂度 O(1)
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// 解法2，类似环形结构，结果有问题
func rotate1(nums []int, k int) {
	for i := 0; i < k; i++ {
		temp := nums[i]
		var res int
		for j := i + k; j < len(nums); {
			nums[j], temp = temp, nums[j]
			j += k
			res = temp
		}
		nums[i] = res
	}
	for i := 0; i < k; i++ {

	}
}
