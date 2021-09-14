package dp

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

解：
dp[i]表示以nums[i]结尾的子数组的最大和
*/

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	result := nums[0]
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		result = max(result, dp[i])
	}

	return result
}
