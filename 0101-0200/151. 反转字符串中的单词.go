package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}

func reverseWords(s string) string {
	temp := []byte(s)
	res := []byte{}
	for i := 0; i < len(temp); i++ {
		if temp[i] == ' ' {
			continue
		}
		for i < len(temp) && temp[i] != ' ' {
			res = append(res, temp[i])
			i++
		}
		i--
		res = append(res, ' ')
	}
	res = res[:len(res)-1]
	r := string(res)

	split := strings.Split(r, " ")

	for i := 0; i < len(split)/2; i++ {
		split[i], split[len(split)-i-1] = split[len(split)-i-1], split[i]
	}
	aa := ""
	for i := 0; i < len(split); i++ {
		aa += split[i] + " "
	}

	return aa[:len(aa)-1]
}
