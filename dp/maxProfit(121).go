package dp

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	result := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		minPrice = min(prices[i], minPrice)
		result = max(result, prices[i]-minPrice)
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
