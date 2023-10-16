package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("asd", "zx"))
}

func mergeAlternately(word1 string, word2 string) string {
	str1 := []byte(word1)
	str2 := []byte(word2)

	leng := len(str1)
	if leng > len(str2) {
		leng = len(str2)
	}

	res := make([]byte, len(str1)+len(str2))
	j := 0
	for i := 0; i < leng; i++ {
		res[j] = str1[i]
		j++
		res[j] = str2[i]
		j++
	}

	if leng < len(str1) {
		str1, str2 = str2, str1
	}

	for i := leng; i < len(str2); i++ {
		res[j] = str2[i]
		j++
	}
	return string(res)
}
