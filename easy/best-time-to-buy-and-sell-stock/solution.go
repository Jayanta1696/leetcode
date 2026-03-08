package main

func maxProfit(prices []int) int {
	maxProfit := 0
	buy := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else {
			if maxProfit < prices[i]-buy {
				maxProfit = prices[i] - buy
			}
		}
	}

	return maxProfit
}

