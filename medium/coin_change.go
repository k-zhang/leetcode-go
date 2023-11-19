package medium

func coinChange(coins []int, amount int) int {
	maxAmount := amount + 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = maxAmount
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}
