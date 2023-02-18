package goLeetCode

func ClimbStairs(n int) int {
	dp := []int{1, 1}
	for i := 2; i < n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n-1]
}
