package easy

import "math"

func isSubSequence(s string, t string) bool {
	var dp [][]int

	dp = make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 0; i <= len(s); i++ {
		dp[i][0] = 0
	}

	for i := 0; i <= len(t); i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Max(float64(dp[i][j-1]), float64(dp[i-1][j])))
			}
		}
	}

	return len(s) == dp[len(s)][len(t)]
}
