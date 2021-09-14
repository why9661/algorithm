package dp

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	result := 1

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}

	return result
}
