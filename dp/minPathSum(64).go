package dp

/*
最小路径和
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
*/

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))

	for i, v := range grid {
		dp[i] = make([]int, len(v))
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	i := len(dp) - 1
	j := len(dp[i]) - 1

	return dp[i][j]
}
