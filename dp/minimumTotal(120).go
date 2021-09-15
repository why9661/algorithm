package dp

/*
三角形最小路径和

说明：
dp[i][j]表示到第i行第j列元素的最小路径和
*/

func minimumTotal(triangle [][]int) int {
	if len(triangle) < 2 {
		return triangle[0][0]
	}

	dp := make([][]int, len(triangle))

	for i, v := range triangle {
		dp[i] = make([]int, len(v))
	}

	dp[0][0] = triangle[0][0]
	dp[1][0] = triangle[0][0] + triangle[1][0]
	dp[1][1] = triangle[0][0] + triangle[1][1]

	for i := 2; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
				continue
			}
			if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
	}

	oLength := len(dp) - 1
	result := dp[oLength][0]
	for i := 0; i < len(dp[oLength]); i++ {
		result = min(result, dp[oLength][i])
	}

	return result
}
