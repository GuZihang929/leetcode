package main

import "strings"

func main() {

}
func reverseVowels(s string) string {
	temp := []int{}
	t := []byte(s)
	for i := 0; i < len(s); i++ {
		if !strings.Contains("aeiouAEIOU", string(t[i])) {
			temp = append(temp, i)
		}
	}

	for i := 0; i < len(temp)/2; i++ {
		t[temp[i]], t[temp[len(temp)-1-i]] = t[temp[len(temp)-1-i]], t[temp[i]]
	}

	return string(t)
}
