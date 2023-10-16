package main

import "strings"

func main() {

}

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	res := ""
	for i := 0; i < len(strs); i++ {
		temp := ""
		for j := len(strs[i]) - 1; j > -1; j-- {
			temp += string(strs[i][j])
		}
		temp += " "
		res += temp
	}
	return res[:len(res)-1]
}
