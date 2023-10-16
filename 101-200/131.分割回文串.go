package main

import "fmt"

func main() {
	fmt.Println(partition("aab"))
}

func partition(s string) [][]string {
	temp, res := []string{}, &[][]string{}
	dfs131(s, temp, res, 0)
	return *res
}

func dfs131(s string, temp []string, res *[][]string, index int) {
	start, end := index, len(s)
	if start == end {
		b := make([]string, len(temp))
		copy(b, temp)
		*res = append(*res, b)
		return
	}

	for i := start; i < end; i++ {
		if IsHuiW(s[start : i+1]) {
			dfs131(s, append(temp, s[start:i+1]), res, i+1)
		}
	}

}

func IsHuiW(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
