package main

func main() {

}

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, len(prices))
	}

	dp[0][0] = 0

	for i := 0; i < len(dp); i++ {

		for j := 0; j < len(dp); j++ {
			if prices[j]-prices[i] > dp[i][j-1] {
				dp[i][j] = prices[j] - prices[i]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	for i := 1; i < len(dp); i++ {
		if dp[len(dp)-1][i] < dp[len(dp)-1][i-1] {
			dp[len(dp)-1][i] = dp[len(dp)-1][i-1]
		}
	}

	return dp[len(dp)-1][len(dp)-1]
}
