package buysell

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minPrice := prices[0]
	mxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			mxProfit = max(mxProfit, prices[i]-minPrice)
		}
	}

	return mxProfit
}
