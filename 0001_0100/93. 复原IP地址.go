package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(restoreIpAddresses("25525511135"))
}

func restoreIpAddresses(s string) []string {

	res := []string{}
	dfs93(s, []int{}, &res, 0)
	return res
}

func dfs93(s string, ip []int, res *[]string, index int) {

	if index == len(s) {
		if len(ip) == 4 {
			*res = append(*res, getString(ip))
		}
		return

	}

	if index == 0 {
		num, _ := strconv.Atoi(string(s[0]))
		ip = append(ip, num)
		dfs93(s, ip, res, index+1)
	} else {
		num, _ := strconv.Atoi(string(s[index]))
		next := ip[len(ip)-1]*10 + num
		if next <= 255 && ip[len(ip)-1] != 0 {
			ip[len(ip)-1] = next
			dfs93(s, ip, res, index+1)
			ip[len(ip)-1] /= 10
		}
		if len(ip) < 4 {
			ip = append(ip, num)
			dfs93(s, ip, res, index+1)
			ip = ip[:len(ip)-1]
		}
	}

}

func getString(ip []int) string {
	res := strconv.Itoa(ip[0])
	for i := 1; i < len(ip); i++ {
		res += "." + strconv.Itoa(ip[i])
	}
	return res
}
