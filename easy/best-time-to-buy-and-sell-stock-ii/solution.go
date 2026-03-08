package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	buy := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if buy < price {
			maxProfit += price - buy
		}
		buy = price
	}

	return maxProfit
}

