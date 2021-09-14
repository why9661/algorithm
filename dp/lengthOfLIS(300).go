package dp

/*
最长上升子序列
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

解：
dp[i]表示以nums[i]结尾的最长上升子序列
*/

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
