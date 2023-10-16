package main

func main() {

}

func maxVowels(s string, k int) int {
	temp := []byte(s)
	res := 0
	for i := 0; i < k; i++ {
		if temp[i] == 'a' || temp[i] == 'e' || temp[i] == 'i' || temp[i] == 'o' || temp[i] == 'u' {
			res++
		}
	}
	max := res
	for i := 1; i < len(s)-k+1; i++ {
		if temp[i-1] == 'a' || temp[i-1] == 'e' || temp[i-1] == 'i' || temp[i-1] == 'o' || temp[i-1] == 'u' {
			res--
		}
		if temp[i+k-1] == 'a' || temp[i+k-1] == 'e' || temp[i+k-1] == 'i' || temp[i+k-1] == 'o' || temp[i+k-1] == 'u' {
			res++
		}
		if max < res {
			max = res
		}

	}
	return max
}
