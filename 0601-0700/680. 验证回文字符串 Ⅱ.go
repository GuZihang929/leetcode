package main

import "fmt"

func main() {

	fmt.Println(validPalindrome("a"))
}

func validPalindrome(s string) bool {

	return left_del(s) || right_del(s)
}

func left_del(s string) bool {
	temp := 0
	for i, j := 0, len(s)-1; j > i; {
		if s[i] != s[j] {
			if temp == 1 {
				return false
			}
			j++
			temp++
		}
		i++
		j--
	}
	return true
}

func right_del(s string) bool {
	temp := 0
	for i, j := 0, len(s)-1; j > i; {
		if s[i] != s[j] {
			if temp == 1 {
				return false
			}
			i--
			temp++
		}
		i++
		j--
	}
	return true
}
