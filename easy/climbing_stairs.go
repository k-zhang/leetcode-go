package easy

func climbingStairs(n int) int {
	if n == 1 {
		return 1
	}

	dp := []int{1, 2}

	for i := 2; i < n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}

	return dp[n-1]
}
