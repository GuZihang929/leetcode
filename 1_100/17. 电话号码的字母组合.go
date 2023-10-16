package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("234234"))
}

//给定一个仅包含数2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

func letterCombinations(digits string) []string {

	arr = make(map[string][]string)
	leng = len(digits)
	if leng == 0 {
		return []string{}
	}
	arr["2"] = []string{"a", "b", "c"}
	arr["3"] = []string{"d", "e", "f"}
	arr["4"] = []string{"g", "h", "i"}
	arr["6"] = []string{"m", "n", "o"}
	arr["7"] = []string{"p", "q", "r", "s"}
	arr["8"] = []string{"t", "u", "v"}
	arr["9"] = []string{"w", "x", "y", "z"}
	arr["5"] = []string{"j", "k", "l"}
	c, res := "", &[]string{}
	dfs17(digits, c, res)

	return *res

}

var arr map[string][]string
var leng int

func dfs17(digits string, c string, res *[]string) {
	if len(c) == leng {
		*res = append(*res, c)
		return
	}

	for _, i2 := range arr[digits[0:1]] {
		c += i2
		dfs17(digits[1:], c, res)
		c = c[:len(c)-1]
	}
}
