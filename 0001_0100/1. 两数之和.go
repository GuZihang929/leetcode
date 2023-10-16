package main

import "fmt"

func main() {

	s := "bbbbb"

	lengthOfLongestSubstring(s)
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	temp := 0
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			i = m[s[i]]
			if temp > max {
				max = temp
			}
			temp = 0
			m = map[byte]int{}
		} else {
			m[s[i]] = i
			temp++
		}
		fmt.Println(m)
	}
	if temp > max {
		max = temp
	}
	return max
}
