package main

import "fmt"

func main() {
	fmt.Println(repeatedSubstringPattern("bb"))
}

func repeatedSubstringPattern(s string) bool {

	for i := 1; i < len(s); i++ {
		temp := s[:i]
		res := temp
		for len(res) < len(s) {
			if res != s[:len(res)] {
				break
			}
			res += temp
		}
		if res == s {
			return true
		}
	}
	return false
}
