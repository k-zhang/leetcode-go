package easy

func bestTimeToBuySellStock(prices []int) int {
	var maxProfit = 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit = max(maxProfit, prices[i]-prices[i-1])
		}

		prices[i] = min(prices[i-1], prices[i])
	}

	return maxProfit
}

func bestTimeToBuySellStock1(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		maxProfit = max(maxProfit, prices[i]-minPrice)
		minPrice = min(prices[i], minPrice)
	}

	return maxProfit
}
