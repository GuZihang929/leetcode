package main

func main() {

}

func sortArrayByParityII(nums []int) []int {
	mid := len(nums) / 2

	num1, num2 := make([]int, mid), make([]int, mid)
	//if num1[0]%2 != 1 {
	//	num1, num2 = num2, num1
	//}
	x, y := 0, 0
	for _, num := range nums {
		if num%2 == 1 {
			num1[x] = num
			x++
		} else {
			num2[y] = num
			y++
		}
	}

	res := make([]int, mid*2)

	for i, j := 0, 0; i < mid*2; i++ {
		res[i] = num2[j]
		i++
		res[i] = num1[j]
		j++

	}
	return res
}
